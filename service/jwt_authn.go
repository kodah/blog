package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string, isUser, isAdmin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name  string `json:"name"`
	User  bool   `json:"user"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    "kodah.blog",
	}
}

func getSecretKey() string {
	cfg := ConfigurationService("")

	return cfg.GetJWTSecret()
}

func (service *jwtServices) GenerateToken(username string, isUser, isAdmin bool) string {
	if !isUser {
		return ""
	}

	claims := &authCustomClaims{
		username,
		isUser,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey)) // encoded
	if err != nil {
		panic(err)
	}

	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
