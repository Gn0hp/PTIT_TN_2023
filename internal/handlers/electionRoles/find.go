package electionRoles

import (
	"PTIT_TN/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h Handler) FindRolesByElectionId(c *gin.Context) {
	electionId := c.Query("election_id")
	if electionId == "" {
		_ = c.Error(errors.New("[Election Role Handler] Invalid election id"))
		return
	}
	electionIdUint, err := strconv.ParseUint(electionId, 10, 64)
	if err != nil {
		_ = c.Error(err)
		return
	}
	roles, err := h.repo.Database().ElectionRole().FindByElectionId(c, electionIdUint)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    roles,
	})
}
