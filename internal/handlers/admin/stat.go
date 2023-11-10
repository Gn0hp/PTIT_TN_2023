package admin

import (
	"PTIT_TN/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// StatByCandidate url: admin/stat-by-candidate/:election_id
func (h *Handler) StatByCandidate(c *gin.Context) {
	electionId, err := strconv.ParseInt(c.Param("election_id"), 10, 64)
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
	var res []*StatByCandidateResponse
	for _, electionRole := range electionRoles {
		candidates, err := h.repo.Database().Candidate().FindCandidateByElectionRoleId(c, electionRole.ID)
		if err != nil {
			h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Find Candidate failed, detail: %v", err))
			_ = c.Error(err)
			continue
		}
		for _, candidate := range candidates {
			res = append(res, &StatByCandidateResponse{
				Candidate:    candidate.Candidate,
				Role:         electionRole.Name,
				ElectionRole: *electionRole,
				RoleElectId:  candidate.RoleElectId,
			})
		}
	}
	for _, stat := range res {
		numVote, err := h.repo.Database().BallotRoleElect().CountVoterByCandidateId(c, stat.RoleElectId)
		if err != nil {
			h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Count Voter failed, detail: %v", err))
			_ = c.Error(err)
			continue
		}
		stat.TotalVote = numVote
		h.logger.Info(fmt.Sprintf("[Voter Handler - ViewCandidate] Stat: %v", stat.TotalVote))
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    res,
	})
}

// StatBallotByCandidate url: /stat-by-candidate/:role_elect_id
func (h *Handler) StatBallotByCandidate(c *gin.Context) {
	roleElectId, err := strconv.ParseInt(c.Param("candidate_id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Invalid role_elect_id failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	resDetail, err := h.repo.Database().BallotRoleElect().StatVoterByRoleElectId(c, uint64(roleElectId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Stat Voter failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    resDetail,
	})
}

func (h *Handler) StatByBallotId(c *gin.Context) {
	ballotId, err := strconv.ParseInt(c.Param("ballot_id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Invalid ballot_id failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	resDetail, err := h.repo.Database().BallotRoleElect().FindByBallotId(c, uint64(ballotId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Stat Voter failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    resDetail,
	})
}
