package models

import (
	"encoding/json"
	"soundpad-backend/dals/entity"
	"strconv"
)

type SoundPadToSound struct {
	SoundPadId json.Number `json:"soundPadId,omitempty"`
	SoundId    json.Number `json:"soundId,omitempty"`
}

func (s *SoundPadToSound) FromEntity(entity *entity.SoundPadToSound) *SoundPadToSound {
	return &SoundPadToSound{
		SoundPadId: json.Number(strconv.FormatInt(entity.SoundPadId, 10)),
		SoundId:    json.Number(strconv.FormatInt(entity.SoundId, 10)),
	}
}

func (s *SoundPadToSound) ToEntity() (*entity.SoundPadToSound, error) {
	soundPadId, err := s.SoundPadId.Int64()
	if err != nil {
		return nil, err
	}
	soundId, err := s.SoundId.Int64()
	if err != nil {
		return nil, err
	}
	return &entity.SoundPadToSound{
		SoundPadId: soundPadId,
		SoundId:    soundId,
	}, nil
}
