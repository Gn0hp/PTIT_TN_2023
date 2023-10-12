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
	var userDto entities.UserDto
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Invalid request body, err: %v", err))
		_ = c.Error(err)
		return
	}
	userDto.Password = utils.HashPassword(userDto.Password)
	if !userDto.IsValid() {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Invalid dto value, dto: %v", userDto))
		_ = c.Error(errors.New("Invalid object value"))
		return
	}
	iEntity, _ := userDto.ToEntity()
	entity := iEntity.(*entities.User)
	err = s.repo.Database().User().Create(entity)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Register] Insert into db fail, err: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}

func (s *Handler) Login(c *gin.Context) {
	var userDto entities.User
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("[User_Handler_Login] Invalid request body, err: %v", err))
		_ = c.Error(err)
		return
	}
	userDto.Password = utils.HashPassword(userDto.Password)
	s.logger.Info(fmt.Sprintf("%v", userDto))

}

func (s *Handler) Logout(c *gin.Context) {
	utils.SetResponse(c, map[string]interface{}{
		"message": "pong",
	})
}
