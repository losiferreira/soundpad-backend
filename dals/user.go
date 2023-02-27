package dals

import (
	"context"
	"github.com/uptrace/bun"
	"log"
	"soundpad-backend/dals/entity"
)

type UserDal struct {
	ctx context.Context
	db  *bun.DB
}

func NewUserDal(
	db *bun.DB,
	ctx context.Context,
) *UserDal {
	return &UserDal{
		db:  db,
		ctx: ctx,
	}
}

func (u *UserDal) CreateUser(user *entity.User) (int64, error) {
	_, err := u.db.NewInsert().Model(user).Exec(u.ctx)
	if err != nil {
		log.Fatalf("Error creating user: %s", err)
		return 0, err
	}
	user, err = u.RetrieveUserByEmail(user.Email)
	return user.Id, nil
}

func (u *UserDal) RetrieveUser(userId string) (*entity.User, error) {
	result := &entity.User{}
	err := u.db.NewSelect().
		Model(result).
		Where("? = ?", bun.Ident("id"), userId).
		Scan(u.ctx)
	return result, err
}

func (u *UserDal) RetrieveUserByEmail(email string) (*entity.User, error) {
	result := &entity.User{}
	err := u.db.NewSelect().
		Model(result).
		Where("? = ?", bun.Ident("email"), email).
		Scan(u.ctx)
	return result, err
}

func (u *UserDal) UpdateUser() {

}

func (u *UserDal) DeleteUser() {

}
