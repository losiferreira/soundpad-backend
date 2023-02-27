package use_cases

import (
	"database/sql"
	"github.com/markbates/goth"
	"log"
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

func (a *AuthUseCase) SignInOrSignUp(gothUser goth.User) (*entity.User, error) {

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
			return nil, err
		}
		soundPadUser.Id = id
		log.Println("New user registered.")
		return soundPadUser, err
	}
	if err != nil {
		log.Fatalf("Error trying to retrieve user by email: %s", err)
		return nil, err
	}
	return soundPadUser, nil
}
