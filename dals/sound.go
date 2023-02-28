package dals

import (
	"context"
	"github.com/uptrace/bun"
	"soundpad-backend/dals/entity"
)

type SoundDal struct {
	ctx context.Context
	db  *bun.DB
}

func NewSoundDal(
	db *bun.DB,
	ctx context.Context,
) *SoundDal {
	return &SoundDal{
		db:  db,
		ctx: ctx,
	}
}

func (s *SoundDal) CreateSound(sound *entity.Sound) (int64, error) {
	_, err := s.db.
		NewInsert().
		Model(sound).
		Returning("id").
		Exec(s.ctx)
	if err != nil {
		return -1, err
	}
	return sound.Id, nil
}

func (s *SoundDal) RetrieveSound(soundId int) (*entity.Sound, error) {
	result := &entity.Sound{}
	err := s.db.NewSelect().
		Model(result).
		Where("? = ?", bun.Ident("id"), soundId).
		Scan(s.ctx)
	return result, err
}

func (s *SoundDal) UpdateSound(sound *entity.Sound) error {
	_, err := s.db.
		NewUpdate().
		Model(sound).
		WherePK().
		Exec(s.ctx)
	return err
}

func (s *SoundDal) DeleteSound(id int) error {
	_, err := s.db.
		NewDelete().
		Model(&entity.Sound{}).
		Where("id = ?", id).
		Exec(s.ctx)
	return err
}
