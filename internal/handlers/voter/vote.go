package voter

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/repositories/candidate"
	"PTIT_TN/pkg"
	"PTIT_TN/pkg/rabbitMQ"
	"PTIT_TN/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (h *Handler) Vote(c *gin.Context) {
	var req VoteRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - Vote] Bind Vote failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	electionResult, err := h.repo.Database().ElectionResult().FindByElectionId(c, req.ElectionId)
	if err != nil || electionResult == nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - Vote] Find Election Result failed, detail: %v", err))
		_ = c.Error(err)
		return
	}

	// Create ballot then ballotRoleElect
	ballotDto := entities.BallotDto{
		Ballot: entities.Ballot{
			VoterId:  req.VoterId,
			ResultId: electionResult.ID,
			Note:     time.Now().String(),
		},
	}
	ballotExisted, _ := h.repo.Database().Ballot().FindByVoterId(c, req.VoterId, req.ElectionId)
	if ballotExisted.ID != 0 {
		h.logger.Error(fmt.Sprintf("[Voter Handler - Vote] Voter already voted"))
		_ = c.Error(fmt.Errorf("voter already voted"))
		return
	}
	iEntity, _ := ballotDto.ToEntity()
	entity := iEntity.(*entities.Ballot)
	err = h.repo.Database().Ballot().Create(c, entity)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - Vote] Create Ballot failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	var ballotRoleElects []*entities.BallotRoleElect
	var choiceValues []int64
	for _, roleElect := range req.RoleElects {
		ballotRoleElectDto := entities.BallotRoleElectDto{
			BallotRoleElect: entities.BallotRoleElect{
				BallotId:    entity.ID,
				RoleElectId: roleElect.RoleElectId,
				Note:        "",
			},
		}
		iEntity, _ := ballotRoleElectDto.ToEntity()
		entity := iEntity.(*entities.BallotRoleElect)
		ballotRoleElects = append(ballotRoleElects, entity)
		choiceValues = append(choiceValues, int64(roleElect.RoleElectId))
	}
	err = h.repo.Database().BallotRoleElect().CreateInBatches(c, ballotRoleElects)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - Vote] Create Ballot Role Elect failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	election, err := h.repo.Database().Election().FindById(c, req.ElectionId)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - Vote] Find Election failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	h.rabbitMQ.Send(rabbitMQ.RabbitRequest{
		Type: pkg.TxTypeVote,
		// electionId, voterId (string), choice value (id)
		Data: map[string]interface{}{
			"election_id":  election.ApiObjectId,
			"voter_id":     strconv.FormatUint(req.VoterId, 10),
			"choice_value": choiceValues,
		},
	})
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) RegisterCandidate(c *gin.Context) {
	// TODO: Check election time
	var req RegisterCandidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - RegisterCandidate] Bind Candidate failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	candidateDto := entities.CandidateDto{
		Candidate: entities.Candidate{
			User:            req.User,
			CccdID:          req.User.CccdID,
			Party:           req.Party,
			Patrimony:       req.Patrimony,
			CandidateStatus: entities.CANDIDATE_STATUS_PENDING,
		},
	}
	candidateDto.User.CccdID = ""
	isRegistered, _ := h.repo.Database().Candidate().CheckRegistered(c, req.User.CccdID, req.Role)
	if isRegistered {
		h.logger.Error(fmt.Sprintf("[Voter Handler - RegisterCandidate] Candidate already registered"))
		_ = c.Error(fmt.Errorf("candidate already registered"))
		return
	}
	roleId := req.Role
	iEntity, err := candidateDto.ToEntity()
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - RegisterCandidate] Convert Candidate to entity failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	entity := iEntity.(*entities.Candidate)
	err = h.repo.Database().Candidate().Create(c, entity)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - RegisterCandidate] Register Candidate failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	roleElect := entities.RoleElectDto{
		RoleElect: entities.RoleElect{
			Name:           fmt.Sprintf("%s %d", entity.FullName, roleId),
			CandidateId:    entity.ID,
			ElectionRoleId: roleId,
		},
	}
	iEntity, err = roleElect.ToEntity()
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - RegisterCandidate] Convert RoleElect to entity failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	entityRoleElect := iEntity.(*entities.RoleElect)
	err = h.repo.Database().RoleElect().Create(c, entityRoleElect)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - RegisterCandidate] Register RoleElect failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) ViewCandidate(c *gin.Context) {
	electionId, err := strconv.ParseInt(c.Query("election_id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Invalid election_id failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	electionRoles, err := h.repo.Database().ElectionRole().FindByElectionId(c, uint64(electionId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Find Election Role failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	res := ViewCandidateResponse{
		MapCandidate: make(map[string][]*candidate.ViewCandidateDbFind),
		SupportedMap: make(map[string]*entities.ElectionRole),
	}
	for _, electionRole := range electionRoles {
		candidates, err := h.repo.Database().Candidate().FindCandidateByElectionRoleId(c, electionRole.ID)
		if err != nil {
			h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Find Candidate failed, detail: %v", err))
			_ = c.Error(err)
			continue
		}
		res.MapCandidate[electionRole.Name] = candidates
		res.SupportedMap[electionRole.Name] = electionRole
	}

	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    res,
	})
}
