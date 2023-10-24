package admin

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/auth"
	"PTIT_TN/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AdminLogin(c *gin.Context) {
	var adminDto entities.AdminDto
	err := c.ShouldBindJSON(&adminDto)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin_Handler_Login] Invalid request body, err: %v", err))
		_ = c.Error(err)
		return
	}
	iAdmin, err := adminDto.ToEntity()
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin_Handler_Login] Cannot convert to Entity, err: %v", err))
		_ = c.Error(err)
		return
	}
	admin := iAdmin.(*entities.Admin)
	adminExisted, err := h.repo.Database().Admin().FindByUsername(c, admin.Username)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin_Handler_Login] Find Admin account failed, err: %v", err))
		_ = c.Error(err)
		return
	}
	if adminExisted.ID == 0 {
		h.logger.Error(fmt.Sprintf("[Admin_Handler_Login] User not found, err: %v", err))
		_ = c.Error(errors.New("user not found"))
		return
	}
	if adminExisted.Password != admin.Password {
		h.logger.Error(fmt.Sprintf("[Admin_Handler_Login] Password not match, err: %v", err))
		_ = c.Error(errors.New("incorrect password"))
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    adminExisted,
		"access_token": auth.New().GenerateAccessToken(adminExisted.Username,
			fmt.Sprintf("%d", adminExisted.ID),
			c),
	})
}
