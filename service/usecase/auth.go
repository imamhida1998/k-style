package usecase

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Auth interface {
	GenerateTokenJWT(email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

var SecretKey = []byte("crud-testing")

func (s *JWTService) GenerateTokenJWT(email string) (string, error) {
	claim := jwt.MapClaims{}
	claim["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	SignToken, err := token.SignedString(SecretKey)
	if err != nil {
		return SignToken, err
	}

	return SignToken, nil

}

func (s *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Token_Invalid")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
