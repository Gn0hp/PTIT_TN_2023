package election

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg"
	"PTIT_TN/pkg/rabbitMQ"
	"PTIT_TN/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CreateElection(c *gin.Context) {
	// TODO: Create New Election
	var createElectionBody CreateElectionBody
	if err := c.ShouldBindJSON(&createElectionBody); err != nil {
		_ = c.Error(err)
		return
	}
	createElectionBody.Election.Status = entities.ELECTION_STATUS_REGISTERING
	electionDto := entities.ElectionDto{
		Election: createElectionBody.Election,
	}
	rolesDto := createElectionBody.Roles
	if !electionDto.IsValid() || !entities.IsAllElectionRoleValid(rolesDto) {
		_ = c.Error(errors.New("[Election Handler] Invalid object value"))
		return
	}
	iEntity, _ := electionDto.ToEntity()
	entity := iEntity.(*entities.Election)
	err := h.repo.Database().Election().Create(c, entity)
	if err != nil {
		_ = c.Error(err)
		return
	}
	err = h.createElectionRelated(c, entity, rolesDto)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Create Election Related Error: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}

func (h *Handler) PushBlockchain(c *gin.Context) {
	// Update Election in off-chain db
	// Update ApiObjectID -> manually by calling nonce
	var electionDto entities.ElectionDto
	electionId := c.Param("id")
	if err := c.ShouldBindJSON(&electionDto); err != nil {
		_ = c.Error(err)
		return
	}
	electionDto.Election.Status = entities.ELECTION_STATUS_OPENING
	iEntity, _ := electionDto.ToEntity()
	entity := iEntity.(*entities.Election)
	entity.ID, _ = strconv.ParseUint(electionId, 10, 64)
	if entity.ID == 0 {
		_ = c.Error(errors.New("[Election Handler] Invalid object value"))
		return
	}
	err := h.repo.Database().Election().UpdateById(c, entity)
	if err != nil {
		_ = c.Error(err)
		return
	}
	// counte RoleElect by electionID
	countElectionRoles, err := h.repo.Database().Election().CountErsByElectionId(c, entity.ID)
	h.mqService.Send(rabbitMQ.RabbitRequest{
		Type: pkg.TxTypeCreateElection,
		Data: CreateElectionReq{
			ID:           entity.ID,
			StartDate:    uint64(entity.DateStartElecting),
			Duration:     uint64(entity.Duration),
			NumCandidate: countElectionRoles, // get count RoleElect in db
		},
	})
	nonce := GetNonceBlockchain(h)
	// validate nonce
	if nonce == nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Get Nonce failed"))
	}
	// convert nonce from to uint64
	nonceUint64, ok := nonce.(uint64)
	if !ok {
		h.logger.Error(fmt.Sprintf("[Election Handler] Convert Nonce failed"))
	}
	entity.ApiObjectId = nonceUint64
	err = h.repo.Database().Election().UpdateById(c, entity)
	if err != nil {
		_ = c.Error(err)
		return
	}
	// Update ApiObjectID

	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}
func (h *Handler) CloseElection(c *gin.Context) {
	electionId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Invalid election_id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	election, err := h.repo.Database().Election().FindById(c, electionId)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Find election failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	election.Status = entities.ELECTION_STATUS_FINISHED
	err = h.repo.Database().Election().UpdateById(c, election)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Update election failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	// TODO: Update all candidate_status to INACTIVE
	err = h.repo.Database().Candidate().DeactivateElectionCandidate(c, electionId)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Deactivate election candidate failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    election,
	})
}

func (h *Handler) createElectionRelated(c *gin.Context, election *entities.Election, rolesDto []entities.ElectionRole) error {
	//ElectionCandidate, Election Result,
	electionCandidateDto := entities.ElectionCandidateDto{
		ElectionCandidate: entities.ElectionCandidate{
			ElectionId: election.ID,
			Note:       "",
		},
	}
	iEntity, _ := electionCandidateDto.ToEntity()
	entity := iEntity.(*entities.ElectionCandidate)
	if err := h.repo.Database().ElectionCandidate().Create(c, entity); err != nil {
		_ = c.Error(err)
		return err
	}
	// Create Election Role
	iRolesEntity, _ := entities.ElectionRoleToEntitiesWithEcId(rolesDto, entity.ID)
	if err := h.repo.Database().ElectionRole().BulkInsert(c, iRolesEntity); err != nil {
		_ = c.Error(err)
		return err
	}
	electionResultDto := entities.ElectionResultDto{
		ElectionResult: entities.ElectionResult{
			ElectionCandidateId: entity.ID,
			Note:                "",
		},
	}
	iEntity1, _ := electionResultDto.ToEntity()
	entity1 := iEntity1.(*entities.ElectionResult)
	if err := h.repo.Database().ElectionResult().Create(c, entity1); err != nil {
		_ = c.Error(err)
		return err
	}
	return nil
}

func GetNonceBlockchain(h *Handler) interface{} {
	resultChan, errChan := make(chan rabbitMQ.RabbitRequest), make(chan error)
	h.mqService.Send(rabbitMQ.RabbitRequest{
		Type: pkg.TxTypeNonce,
		Data: nil,
	})
	go h.mqService.Receive(resultChan, errChan)
	select {
	case res := <-resultChan:
		return res
	case err := <-errChan:
		h.logger.Error(fmt.Sprintf("[Election Handler] Get Nonce failed, detail: %v", err))
		return nil
	}
}
