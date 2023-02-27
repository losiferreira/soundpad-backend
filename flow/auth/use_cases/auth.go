package use_cases

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth"
	"log"
	"os"
	"soundpad-backend/dals"
	"soundpad-backend/dals/entity"
)

type AuthUseCase struct {
	dal *dals.UserDal
}

func NewAuthUsecase(
	dal *dals.UserDal,
) *AuthUseCase {
	return &AuthUseCase{
		dal: dal,
	}
}

func (a *AuthUseCase) SignInOrSignUp(gothUser goth.User) (string, error) {

	soundPadUser, err := a.dal.RetrieveUserByEmail(gothUser.Email)
	log.Printf("Soundpad user = %s\n", soundPadUser.Email)

	if soundPadUser == nil || err == sql.ErrNoRows {
		log.Println("User was not registered.")
		soundPadUser = &entity.User{
			Name:  gothUser.Name,
			Email: gothUser.Email,
		}
		id, err := a.dal.CreateUser(soundPadUser)
		if err != nil {
			log.Printf("Error creating user: %s\n", err)
			return "", err
		}
		soundPadUser.Id = id
		log.Println("New user registered.")
	}

	token, err := generateJwt(soundPadUser)

	return token, err
}

func generateJwt(user *entity.User) (string, error) {
	key := []byte(os.Getenv("AUTH_HMAC_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatalf("Could not generate JWT: %s", err)
		return "", err
	}

	log.Printf("JWT = %s", tokenString)
	return tokenString, nil
}
