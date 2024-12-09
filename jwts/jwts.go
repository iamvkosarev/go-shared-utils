package jwts

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JWT struct {
	secret   []byte
	duration time.Duration
}

func NewJWT(secret string, duration time.Duration) JWT {
	return JWT{secret: []byte(secret), duration: duration}
}

func (j JWT) GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.duration).Unix(), // Token expires in 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}
