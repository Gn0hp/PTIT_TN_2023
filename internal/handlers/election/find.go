package election

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/repositories/candidate"
	"PTIT_TN/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ElectionCheckResponse struct {
	HasElection bool   `json:"has_election"`
	ElectionID  uint64 `json:"election_id"`
}

func (h *Handler) CheckElection(c *gin.Context) {

	dataCheck, err := h.repo.Database().Election().FindOpenElection(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	if dataCheck == nil {
		utils.SetResponse(c, map[string]interface{}{
			"success": true,
			"data": ElectionCheckResponse{
				HasElection: false,
				ElectionID:  0,
			},
		})
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data": ElectionCheckResponse{
			HasElection: true,
			ElectionID:  dataCheck.ID,
		},
	})

}

func (h *Handler) ViewResult(c *gin.Context) {
	electionId, err := h.repo.Database().Election().FindLatestElection(c)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Voter Handler - ViewCandidate] Find Latest Election failed, detail: %v", err))
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
		MapCandidate: make(map[string][]*candidate.ViewResultDbFind),
		SupportedMap: make(map[string]*entities.ElectionRole),
	}
	for _, electionRole := range electionRoles {
		candidates, err := h.repo.Database().Candidate().ViewResult(c, electionRole.ID)
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
