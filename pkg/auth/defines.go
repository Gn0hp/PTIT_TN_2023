package auth

import "time"

type Authentication struct {
	AccessToken string `json:"access_token"`
	Success     bool   `json:"success"`
	Error       string `json:"error,omitempty"`
}

const (
	AccessTokenKey        = "access-token"
	IssuerAddressClaimKey = "iss"
	IssuerIdClaimKey      = "jti"
)

const AccessTokenExpiry = time.Minute * 20
