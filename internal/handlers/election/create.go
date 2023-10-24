package election

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateElection(c *gin.Context) {
	// TODO: Create New Election, Election Candidate, Election Result,
	var electionDto entities.ElectionDto
	if err := c.ShouldBindJSON(&electionDto); err != nil {
		_ = c.Error(err)
		return
	}
	if !electionDto.IsValid() {
		_ = c.Error(errors.New("[Election Handler] Invalid object value"))
		return
	}

	iEntity, _ := electionDto.ToEntity()
	entity := iEntity.(*entities.Election)
	if err := h.repo.Database().Election().Create(c, entity); err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}

func (h *Handler) PushBlockchain(c *gin.Context) {
	// TODO: Push RabbitMQ save Blockchain
	// Update Election in off-chain db
	var electionDto entities.ElectionDto
	if err := c.ShouldBindJSON(&electionDto); err != nil {
		_ = c.Error(err)
		return
	}

}
