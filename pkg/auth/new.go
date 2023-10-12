package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"logur.dev/logur"
	"net/http"
	"strings"
	"time"
)

type Authenticator interface {
	GenerateAccessToken(address, id string, c *gin.Context) string
}

type impl struct {
}

func New() Authenticator {
	return impl{}
}

func (i impl) GenerateAccessToken(username, id string, c *gin.Context) string {
	return newJwtService().generateAccessToken(username, id, c, time.Now().Add(AccessTokenExpiry))
}

// Middleware checks if request comes with valid access token
func Middleware(logger logur.LoggerFacade) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !skipTokenCheck(c.Request.URL.Path) {
			at := extractAccessToken(c)
			if at == "" {
				logger.Info("[Auth] no access token, proceed to extract refresh token")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// if access token present
			t, err := newJwtService().validateToken(at, c)
			if err != nil {
				logger.Error(fmt.Sprintf("[Auth] error when verify access token, detail: %v", err))
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Extract token metadata and store its tokenDetails into the same request context
			userDetails, err := extractTokenMetadata(t)
			if err != nil || userDetails == nil {
				logger.Error(fmt.Sprintf("[Auth] error when extract token metadata, detail: %v", err))
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			setUserContext(c, *userDetails)
			c.Next()
		}
	}
}

func extractAccessToken(c *gin.Context) string {
	// Get token from Authorization header
	bearerToken := c.GetHeader("Authorization")
	// Split the token
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 && len(strArr[1]) > 0 {
		return strArr[1]
	}
	return ""
}

// extractTokenMetadata extracts metadata from token
func extractTokenMetadata(token *jwt.Token) (*UserDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrorMapClaimNotFound
	}

	userID, ok := claims[IssuerIdClaimKey]
	if !ok {
		return nil, ErrorUserIDNotFound
	}

	username, ok := claims[IssuerAddressClaimKey]
	if !ok {
		return nil, ErrorUserAddressNotFound
	}

	return &UserDetails{
		UserID:   cast.ToInt64(userID),
		Username: username.(string),
	}, nil
}

func skipTokenCheck(uri string) bool {
	skipPath := []string{
		"/api/v1/liveness",
		"/favicon.ico",
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/api/v1/debug/pprof",
		"/swagger",
	}
	for _, s := range skipPath {
		if strings.Contains(uri, s) {
			return true
		}
	}
	return false
}
