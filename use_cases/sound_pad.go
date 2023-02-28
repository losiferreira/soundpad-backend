package use_cases

import (
	"log"
	"soundpad-backend/dals"
	"soundpad-backend/dals/entity"
	"soundpad-backend/handlers/models"
)

type SoundPadUseCase struct {
	dal *dals.SoundPadDal
}

func NewSoundPadUseCase(
	dal *dals.SoundPadDal,
) *SoundPadUseCase {
	return &SoundPadUseCase{
		dal: dal,
	}
}

func (s *SoundPadUseCase) CreateSoundPad(
	ownerId int64,
	soundPad *models.SoundPad,
) error {
	soundPadEntity := &entity.SoundPad{
		Name:    soundPad.Name,
		OwnerId: ownerId,
	}
	_, err := s.dal.CreateSoundPad(soundPadEntity)
	if err != nil {
		log.Printf("Could not create soundPad: %s", err)
		return err
	}

	return nil
}

func (s *SoundPadUseCase) RetrieveSoundPad(
	soundPadId int64,
) (*models.SoundPad, error) {

	soundPadEntity, err := s.dal.RetrieveSoundPad(soundPadId)
	if soundPadEntity == nil || err != nil {
		log.Printf("Could not retrieve soundPad: %s", err)
		return nil, err
	}

	return (&models.SoundPad{}).FromEntity(soundPadEntity), nil
}

func (s *SoundPadUseCase) UpdateSoundPad(
	soundPadModel *models.SoundPad,
) (*models.SoundPad, error) {

	id, _ := soundPadModel.Id.Int64()
	soundPadEntity, err := s.dal.RetrieveSoundPad(id)
	if soundPadEntity == nil || err != nil {
		log.Printf("Could not retrieve soundPadModel: %s", err)
		return nil, err
	}

	updatedSounds := make([]*entity.Sound, len(soundPadModel.Sounds))
	for i := range soundPadModel.Sounds {
		updatedSound, err := soundPadModel.Sounds[i].ToEntity()
		if err != nil {
			log.Printf("Could not convert sound to entity: %s", err)
			return nil, err
		}
		updatedSounds[i] = updatedSound
	}

	soundPadEntity.Name = soundPadModel.Name
	soundPadEntity.Sounds = updatedSounds

	soundPadEntity, err = s.dal.UpdateSoundPad(soundPadEntity)
	if err != nil {
		log.Printf("Could not create soundPadModel: %s", err)
		return nil, err
	}

	return (&models.SoundPad{}).FromEntity(soundPadEntity), nil
}

func (s *SoundPadUseCase) DeleteSoundPad(soundPadId int64) error {
	soundPadEntity, err := s.dal.RetrieveSoundPad(soundPadId)
	if soundPadEntity == nil || err != nil {
		log.Printf("Could not retrieve soundPad: %s", err)
		return err
	}

	err = s.dal.DeleteSoundPad(soundPadId)
	if err != nil {
		log.Printf("Could not delete soundPad: %s", err)
		return err
	}

	return nil
}
