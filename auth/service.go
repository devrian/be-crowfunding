package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// Service is ...
type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

// NewService is ...
func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	payloadClaim := jwt.MapClaims{}
	payloadClaim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadClaim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(jwtToken *jwt.Token) (interface{}, error) {
		_, merged := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !merged {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
