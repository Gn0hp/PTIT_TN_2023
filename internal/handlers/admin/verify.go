package admin

import (
	"PTIT_TN/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) VerifyVoter(c *gin.Context) {

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
