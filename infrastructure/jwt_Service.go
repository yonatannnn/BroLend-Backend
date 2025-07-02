package infrastructure

import (
	"brolend/domain"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(user domain.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (js *jwtService) GenerateToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"user_id":    user.UserID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})

	jwtToken, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (js *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	claims := &jwt.MapClaims{}
	tokenObj, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !tokenObj.Valid {
		return nil, errors.New("invalid token")
	}
	return tokenObj, nil
}
