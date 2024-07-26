package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(id int) (string, error)
	Validate(token string) bool
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "secret-key",
		issuer:    "product-api",
	}
}

type Claim struct {
	Sum int `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id int) (string, error) {
	claim := &Claim{
		Sum: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) Validate(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		fmt.Println("Token parsing error:", err)
		return false
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		fmt.Println("Token is valid for user ID:", claims.Sum)
		return true
	} else {
		fmt.Println("Invalid token")
		return false
	}
}
