package dals

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"soundpad-backend/dals/entity"
)

type SoundPadSoundsDal struct {
	ctx context.Context
	db  *bun.DB
}

func NewSoundPadSoundsDal(
	db *bun.DB,
	ctx context.Context,
) *SoundPadSoundsDal {
	return &SoundPadSoundsDal{
		db:  db,
		ctx: ctx,
	}
}

func (s *SoundPadSoundsDal) CreateSoundPadSounds(soundPadSounds *entity.SoundPadToSound) (int64, error) {
	_, err := s.db.NewInsert().
		Model(soundPadSounds).
		Returning("*").
		Exec(s.ctx)
	if err != nil {
		return -1, err
	}
	return soundPadSounds.SoundPadId, nil
}

func (s *SoundPadSoundsDal) RetrieveSoundPadToSound(soundPadToSound *entity.SoundPadToSound) (*entity.SoundPadToSound, error) {
	err := s.db.NewSelect().
		Model(soundPadToSound).
		Scan(s.ctx)
	return soundPadToSound, err
}

func (s *SoundPadSoundsDal) SoundAlreadyAdded(
	soundPadSoundsToSound *entity.SoundPadToSound,
) (bool, error) {

	_, err := s.db.
		NewSelect().
		Model(soundPadSoundsToSound).
		Exec(s.ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	return true, nil
}

func (s *SoundPadSoundsDal) AddSound(
	soundPadSoundsToSound *entity.SoundPadToSound,
) (*entity.SoundPadToSound, error) {

	_, err := s.db.
		NewInsert().
		Model(soundPadSoundsToSound).
		Returning("*").
		Exec(s.ctx)
	if err != nil {
		return nil, err
	}

	return soundPadSoundsToSound, err
}

func (s *SoundPadSoundsDal) RemoveSound(
	soundPadSoundsToSound *entity.SoundPadToSound,
) error {

	_, err := s.db.
		NewDelete().
		Model(soundPadSoundsToSound).
		Exec(s.ctx)
	if err != nil {
		return err
	}

	return err
}

func (s *SoundPadSoundsDal) DeleteSoundPadSounds(
	soundPadSoundsToSound *entity.SoundPadToSound,
) error {
	_, err := s.db.
		NewDelete().
		Model(soundPadSoundsToSound).
		WherePK().
		Exec(s.ctx)
	return err
}
