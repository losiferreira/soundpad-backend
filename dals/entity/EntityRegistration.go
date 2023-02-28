package entity

import "github.com/uptrace/bun"

type Registration struct {
	db *bun.DB
}

func NewEntityRegistration(
	db *bun.DB,
) *Registration {
	return &Registration{
		db: db,
	}
}

func (r *Registration) Setup() {
	r.db.RegisterModel((*SoundPadToSound)(nil))
}
