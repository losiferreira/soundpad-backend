package providers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"soundpad-backend/dals/entity"
	"strconv"
)

type JwtHandler struct {
}

func NewJwtHandler() *JwtHandler {
	return &JwtHandler{}
}

func (j *JwtHandler) CreateToken(user *entity.User) (string, error) {
	key := []byte(os.Getenv("AUTH_HMAC_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    strconv.FormatInt(user.Id, 10),
		"email": user.Email,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Printf("Could not generate JWT: %s", err)
		return "", err
	}
	log.Printf("token: %s", tokenString)
	return tokenString, nil
}

func (j *JwtHandler) GetToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("AUTH_HMAC_SECRET")), nil
	})
}

func (j *JwtHandler) GetUserIdFromToken(token *jwt.Token) (int64, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idString := claims["id"].(string)
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return -1, fmt.Errorf("error getting token id as int64")
		}
		return id, nil
	} else {
		return -1, fmt.Errorf("error getting token information")
	}
}
