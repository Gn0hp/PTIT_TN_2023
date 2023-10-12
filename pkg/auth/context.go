package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ErrUserNotFoundInCtx = errors.New("user not found in context")
)

const (
	userKey = "user"
)

type UserDetails struct {
	UserID   int64
	Username string
}

// setUserContext sets user details in context
func setUserContext(c *gin.Context, details UserDetails) {
	c.Set(userKey, details)
}

// GetUserFromContext gets user details from context
func GetUserFromContext(c *gin.Context) (UserDetails, error) {
	ud, ok := c.Get(userKey)
	if !ok {
		return UserDetails{}, errors.WithStack(ErrUserNotFoundInCtx)
	}
	return ud.(UserDetails), nil
}
