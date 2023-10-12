package auth

import (
	"crypto"
	"crypto/subtle"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

var secretKey string

func init() {
	genSecret()
	logrus.Info("Init new jwt service")
}

type JwtService struct {
	secretKey     string
	secretKeyByte []byte
}

func genSecret() {
	if len(secretKey) == 0 {
		sha256 := crypto.SHA256.New()
		sha256.Write([]byte(fmt.Sprintf("%v_&_%v", time.Now().Unix(), time.Now().UnixNano())))
		secretKey = string(sha256.Sum(nil))
	}
}

func newJwtService() *JwtService {
	genSecret()
	return &JwtService{
		secretKey:     secretKey,
		secretKeyByte: []byte(secretKey),
	}
}

func (j *JwtService) generateAccessToken(username, id string, c *gin.Context, expiredTime time.Time) string {
	aud := parseClaimString(c)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:   username,
		Audience: aud,
		ExpiresAt: &jwt.NumericDate{
			Time: expiredTime,
		},
		ID: id,
	})
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *JwtService) validateToken(tokenString string, c *gin.Context) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKeyByte, nil
	})
	if err != nil {
		return
	}

	if !token.Valid {
		err = ErrorTokenExpired
		return
	}

	aud := parseClaimString(c)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrorMapClaimNotFound
	}
	for _, str := range aud {
		if !verifyAudience(claims, str, true) {
			err = ErrorTokenAudience
			return
		}
	}
	return token, nil

}
func parseClaimString(c *gin.Context) jwt.ClaimStrings {
	var (
		host      = c.Request.RemoteAddr
		userAgent = c.Request.UserAgent()
	)
	return jwt.ClaimStrings{
		host,
		userAgent,
	}
}

// --- helper

// verifyAudience convert jwt.MapClaims into []string for comparison
func verifyAudience(mapClaims jwt.MapClaims, cmp string, req bool) bool {
	var aud []string
	switch v := mapClaims["aud"].(type) {
	case string:
		aud = append(aud, v)
	case []string:
		aud = v
	case []interface{}:
		for _, a := range v {
			vs, ok := a.(string)
			if !ok {
				return false
			}
			aud = append(aud, vs)
		}
	}
	return verifyAud(aud, cmp, req)
}

// verifyAud verifies the aud claim against the cmp string
// if required is true, then the aud claim must be present or
func verifyAud(aud []string, cmp string, required bool) bool {
	if len(aud) == 0 {
		return !required
	}
	// use a var here to keep constant time compare when looping over a number of claims
	result := false

	var stringClaims string
	for _, a := range aud {
		if subtle.ConstantTimeCompare([]byte(a), []byte(cmp)) != 0 {
			result = true
		}
		stringClaims = stringClaims + a
	}

	// case where "" is sent in one or many aud claims
	if len(stringClaims) == 0 {
		return !required
	}

	return result
}
