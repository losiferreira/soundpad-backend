package use_cases

import (
	"soundpad-backend/dals"
	"soundpad-backend/handlers/models"
)

type SoundPadSoundsUseCase struct {
	dal *dals.SoundPadSoundsDal
}

func NewSoundPadSoundsUseCase(
	dal *dals.SoundPadSoundsDal,
) *SoundPadSoundsUseCase {
	return &SoundPadSoundsUseCase{
		dal: dal,
	}
}

func (s *SoundPadSoundsUseCase) CreateSoundPadSounds(
	soundPadToSound *models.SoundPadToSound,
) error {
	soundPadSoundsEntity, err := soundPadToSound.ToEntity()
	if err != nil {
		return err
	}

	_, err = s.dal.CreateSoundPadSounds(soundPadSoundsEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *SoundPadSoundsUseCase) RetrieveSoundPadSounds(
	soundPadToSound *models.SoundPadToSound,
) (*models.SoundPadToSound, error) {

	soundPadToSoundEntity, err := soundPadToSound.ToEntity()
	soundPadSoundsEntity, err := s.dal.RetrieveSoundPadToSound(soundPadToSoundEntity)
	if soundPadSoundsEntity == nil || err != nil {
		return nil, err
	}

	return (&models.SoundPadToSound{}).FromEntity(soundPadSoundsEntity), nil
}

func (s *SoundPadSoundsUseCase) DeleteSoundPadSounds(
	soundPadToSound *models.SoundPadToSound,
) error {
	soundPadToSoundEntity, err := soundPadToSound.ToEntity()

	err = s.dal.DeleteSoundPadSounds(soundPadToSoundEntity)
	if err != nil {
		return err
	}

	return nil
}
