package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617aj564u564a654u65465aufhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353K554245987uaS?SFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func generateJWTToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
