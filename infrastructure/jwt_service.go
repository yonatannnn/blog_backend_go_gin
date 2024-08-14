package infrastructure

import (
	"blog_api/domain"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(user domain.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) GenerateToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	jwtToken , err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil


}
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	claims := &jwt.MapClaims{}
	tokenObj, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tokenObj.Valid {
		return nil, errors.New("invalid token")
	}
	return tokenObj, nil
}
