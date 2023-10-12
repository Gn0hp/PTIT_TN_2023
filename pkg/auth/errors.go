package auth

import "errors"

var (
	HttpUnauthorizedError = errors.New("invalid credential")

	ErrorUserIDNotFound      = errors.New("user id not found in token claim")
	ErrorUserAddressNotFound = errors.New("user address not found in token claim")
	ErrorMapClaimNotFound    = errors.New("map claim not found in token")
	ErrorTokenExpired        = errors.New("Token is expired")
	ErrorTokenAudience       = errors.New("Token invalid audience")
)

func ErrorToLogFields(key string, err error) map[string]interface{} {
	return map[string]interface{}{
		key: err,
	}
}
