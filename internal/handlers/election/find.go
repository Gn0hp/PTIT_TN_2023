package election

import (
	"PTIT_TN/pkg/utils"
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
