package use_cases

import (
	"database/sql"
	"github.com/markbates/goth"
	"log"
	"soundpad-backend/dals"
	"soundpad-backend/dals/entity"
	"soundpad-backend/providers"
)

type AuthUseCase struct {
	jwt *providers.JwtHandler
	dal *dals.UserDal
}

func NewAuthUsecase(
	jwt *providers.JwtHandler,
	dal *dals.UserDal,
) *AuthUseCase {
	return &AuthUseCase{
		jwt: jwt,
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

	token, err := a.jwt.CreateToken(soundPadUser)

	return token, err
}
