package service

import (
	"tenant-management/cmd/internal/user/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(user *domain.User) string
	ValidateToken(tokenStr string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secret string) JWTService {
	return &jwtService{secretKey: secret}
}

func (j *jwtService) GenerateToken(user *domain.User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(j.secretKey))
	return signed
}

func (j *jwtService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}
