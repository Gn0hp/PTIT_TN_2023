package admin

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) VerifyVoter(c *gin.Context) {
	var voterId struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&voterId)
	if err != nil {
		h.logger.Error("[Admin_Handler_VerifyVoter] Invalid request body")
		_ = c.Error(err)
		return
	}
	voterUpdate := entities.Voter{
		User: entities.User{
			DefaultModel: entities.DefaultModel{
				ID: uint64(voterId.Id),
			},
			Status: entities.REGISTER_STATUS_ACTIVE,
		},
	}
	err = h.repo.Database().Voter().UpdateById(c, &voterUpdate)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) VerifyCandidate(c *gin.Context) {
	var candidateId struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&candidateId)
	if err != nil {
		h.logger.Error("[Admin_Handler_VerifyCandidate] Invalid request body")
		_ = c.Error(err)
		return
	}
	candidateUpdate := entities.Candidate{
		User: entities.User{
			Status: entities.REGISTER_STATUS_ACTIVE,
		},
		CandidateStatus: entities.CANDIDATE_STATUS_VERIFIED,
	}
	err = h.repo.Database().Candidate().UpdateById(c, uint64(candidateId.Id), &candidateUpdate)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) RejectVoter(c *gin.Context) {
	var voterId struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&voterId)
	if err != nil {
		h.logger.Error("[Admin_Handler_RejectVoter] Invalid request body")
		_ = c.Error(err)
		return
	}
	voterUpdate := entities.Voter{
		User: entities.User{
			DefaultModel: entities.DefaultModel{
				ID: uint64(voterId.Id),
			},
			Status: entities.REGISTER_STATUS_REJECTED,
		},
	}
	err = h.repo.Database().Voter().UpdateById(c, &voterUpdate)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}

func (h *Handler) RejectCandidate(c *gin.Context) {
	var candidateId struct {
		Id int `json:"id"`
	}
	err := c.ShouldBindJSON(&candidateId)
	if err != nil {
		h.logger.Error("[Admin_Handler_RejectCandidate] Invalid request body")
		_ = c.Error(err)
		return
	}
	candidateUpdate := entities.Candidate{
		User: entities.User{
			Status: entities.REGISTER_STATUS_REJECTED,
		},
		CandidateStatus: entities.CANDIDATE_STATUS_VERIFIED,
	}
	err = h.repo.Database().Candidate().UpdateById(c, uint64(candidateId.Id), &candidateUpdate)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) GetPendingUserList(c *gin.Context) {
	voters, err := h.repo.Database().Voter().GetPendingVoters(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	candidates, err := h.repo.Database().Candidate().GetPendingCandidate(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"voters":     voters,
			"candidates": candidates,
		},
	})
}
