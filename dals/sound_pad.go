package dals

import (
	"context"
	"github.com/uptrace/bun"
	"log"
	"soundpad-backend/dals/entity"
)

type SoundPadDal struct {
	ctx context.Context
	db  *bun.DB
}

func NewSoundPadDal(
	db *bun.DB,
	ctx context.Context,
) *SoundPadDal {
	return &SoundPadDal{
		db:  db,
		ctx: ctx,
	}
}

func (s *SoundPadDal) CreateSoundPad(soundPad *entity.SoundPad) (int64, error) {
	_, err := s.db.NewInsert().Model(soundPad).Returning("id").Exec(s.ctx)
	if err != nil {
		log.Printf("Error creating soundPad: %s", err)
		return -1, err
	}
	return soundPad.Id, nil
}

func (s *SoundPadDal) RetrieveSoundPad(soundPadId int64) (*entity.SoundPad, error) {
	result := &entity.SoundPad{}
	err := s.db.NewSelect().
		Model(result).
		Where("? = ?", bun.Ident("id"), soundPadId).
		Scan(s.ctx)
	return result, err
}

func (s *SoundPadDal) UpdateSoundPad(soundPad *entity.SoundPad) (*entity.SoundPad, error) {
	_, err := s.db.
		NewUpdate().
		Model(soundPad).
		WherePK().
		Returning("*").
		Exec(s.ctx)
	return soundPad, err
}

func (s *SoundPadDal) DeleteSoundPad(id int64) error {
	_, err := s.db.NewDelete().Where("id = ?", id).Exec(s.ctx)
	return err
}
