package jwts

import (
	"fmt"
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

type JWTValidator struct {
	secret []byte
}

func (j JWTValidator) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return j.secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func NewJWTValidator(secret string) *JWTValidator {
	return &JWTValidator{secret: []byte(secret),}
}

func (j JWT) GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.duration).Unix(), // Token expires in 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}
