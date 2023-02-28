package dals

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
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
	_, err := s.db.
		NewInsert().
		Model(soundPad).
		Returning("id").
		Exec(s.ctx)
	if err != nil {
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
	if err != nil {
		return nil, err
	}

	return soundPad, err
}

func (s *SoundPadDal) SoundAlreadyAdded(
	soundPadToSound *entity.SoundPadToSound,
) (bool, error) {

	_, err := s.db.
		NewSelect().
		Model(soundPadToSound).
		Exec(s.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return true, nil
}

func (s *SoundPadDal) AddSound(
	soundPadToSound *entity.SoundPadToSound,
) (*entity.SoundPadToSound, error) {

	_, err := s.db.
		NewInsert().
		Model(soundPadToSound).
		Returning("*").
		Exec(s.ctx)
	if err != nil {
		return nil, err
	}

	return soundPadToSound, err
}

func (s *SoundPadDal) RemoveSound(
	soundPadToSound *entity.SoundPadToSound,
) error {

	_, err := s.db.
		NewDelete().
		Model(soundPadToSound).
		Exec(s.ctx)
	if err != nil {
		return err
	}

	return err
}

func (s *SoundPadDal) DeleteSoundPad(id int64) error {
	_, err := s.db.
		NewDelete().
		Model(&entity.SoundPad{}).
		Where("id = ?", id).
		Exec(s.ctx)
	return err
}
