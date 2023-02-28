package models

import (
	"encoding/json"
	"soundpad-backend/dals/entity"
	"strconv"
)

type SoundPad struct {
	Id      json.Number `json:"id,omitempty"`
	Name    string      `json:"name,omitempty"`
	OwnerId json.Number `json:"ownerId,omitempty"`

	Sounds []*Sound
}

func (s *SoundPad) FromEntity(entity *entity.SoundPad) *SoundPad {
	sounds := make([]*Sound, len(entity.Sounds))
	for i := range entity.Sounds {
		sounds[i] = (&Sound{}).FromEntity(entity.Sounds[i])
	}

	return &SoundPad{
		Id:      json.Number(strconv.FormatInt(entity.Id, 10)),
		Name:    entity.Name,
		OwnerId: json.Number(strconv.FormatInt(entity.OwnerId, 10)),
		Sounds:  sounds,
	}
}
