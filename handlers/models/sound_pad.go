package models

import (
	"soundpad-backend/dals/entity"
	"strconv"
)

type SoundPad struct {
	Id      string
	Name    string
	OwnerId string

	Sounds []*Sound
}

func (s *SoundPad) FromEntity(entity *entity.SoundPad) *SoundPad {
	sounds := make([]*Sound, len(entity.Sounds))
	for i := range entity.Sounds {
		sounds[i] = (&Sound{}).FromEntity(entity.Sounds[i])
	}

	return &SoundPad{
		Id:      strconv.FormatInt(entity.Id, 10),
		Name:    entity.Name,
		OwnerId: strconv.FormatInt(entity.OwnerId, 10),
		Sounds:  sounds,
	}
}
