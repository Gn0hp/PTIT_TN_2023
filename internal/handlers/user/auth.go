package user

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/auth"
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
	voterDto.Status = entities.REGISTER_STATUS_PENDING
	if !voterDto.IsValid() {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Invalid dto value, dto: %v", voterDto))
		_ = c.Error(errors.New("Invalid object value"))
		return
	}
	iEntity, _ := voterDto.ToEntity()
	entity := iEntity.(*entities.Voter)
	err = s.repo.Database().Voter().Create(c, entity)
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

func (s *Handler) Login(c *gin.Context) {
	var userDto entities.UserDto
	flagVoter := false
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Login] Invalid request body, err: %v", err))
		_ = c.Error(err)
		return
	}
	//userDto.Password = utils.HashPassword(userDto.Password)
	s.logger.Info(fmt.Sprintf("%v", userDto))
	iUser, err := userDto.ToEntity()
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Login] Cannot convert to Entity, err: %v", err))
		_ = c.Error(err)
		return
	}
	entity, _ := iUser.(*entities.User)
	voterExisted, err := s.repo.Database().Voter().FindByCccd(c, entity.CccdID)
	if err != nil || voterExisted == nil || voterExisted.ID == 0 {
		flagVoter = true
	}
	candidateExisted, err := s.repo.Database().Candidate().FindByCccd(c, entity.CccdID)
	if err != nil || candidateExisted == nil || candidateExisted.ID == 0 {
		if flagVoter {
			s.logger.Error(fmt.Sprintf("[User_Login] User not found"))
			_ = c.Error(errors.New("user not found"))
			return
		} else if !utils.CheckHashPassword(entity.Password, voterExisted.Password) {
			s.logger.Error(fmt.Sprintf("[User_Login_Voter] Password not match"))
			_ = c.Error(errors.New("incorrect password"))
			return
		}
		utils.SetResponse(c, map[string]interface{}{
			"success": true,
			"type":    "voter",
			"access_token": auth.New().GenerateAccessToken(
				voterExisted.CccdID,
				fmt.Sprintf("%d", voterExisted.ID), c),
			"data": voterExisted,
		})
		return
	} else {
		if !utils.CheckHashPassword(entity.Password, candidateExisted.Password) {
			s.logger.Error(fmt.Sprintf("[User_Login_Candidate] Password not match"))
			_ = c.Error(errors.New("incorrect password"))
			return
		}
		utils.SetResponse(c, map[string]interface{}{
			"success": true,
			"type":    "candidate",
			"access_token": auth.New().GenerateAccessToken(
				candidateExisted.CccdID,
				fmt.Sprintf("%d", candidateExisted.ID), c),
			"data": candidateExisted,
		})
		return
	}
}

func (s *Handler) Logout(c *gin.Context) {
	utils.SetResponse(c, map[string]interface{}{
		"message": "pong",
	})
}
