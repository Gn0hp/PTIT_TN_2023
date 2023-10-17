package user

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *Handler) Register(c *gin.Context) {
	//TODO implement me
	var voterDto entities.VoterDto
	err := c.ShouldBindJSON(&voterDto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Invalid request body, err: %v", err))
		_ = c.Error(err)
		return
	}
	voterDto.Password = utils.HashPassword(voterDto.Password)
	if !voterDto.IsValid() {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Invalid dto value, dto: %v", voterDto))
		_ = c.Error(errors.New("Invalid object value"))
		return
	}
	iEntity, _ := voterDto.ToEntity()
	entity := iEntity.(*entities.Voter)
	err = s.repo.Database().Voter().Create(entity)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Insert into db fail, err: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}

func (s *Handler) RegisterCandidate(c *gin.Context) {
	var candidateDto entities.CandidateDto
	err := c.ShouldBindJSON(&candidateDto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_RegisterCandidate] Invalid request body, err: %v", err))
		_ = c.Error(errors.New("Invalid object value"))
		return
	}

}
