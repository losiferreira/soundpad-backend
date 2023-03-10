package models

import (
	"encoding/json"
	"soundpad-backend/dals/entity"
	"strconv"
)

type Sound struct {
	Id       json.Number `json:"id,string,omitempty"`
	Name     string      `json:"name,omitempty"`
	FileName string      `json:"fileName,omitempty"`
}

func (s *Sound) FromEntity(sound *entity.Sound) *Sound {
	s.Id = json.Number(strconv.FormatInt(sound.Id, 10))
	s.Name = sound.FileName
	s.FileName = sound.FileName
	return s
}

func (s *Sound) ToEntity() (*entity.Sound, error) {
	id, err := s.Id.Int64()
	if err != nil {
		return nil, err
	}
	return &entity.Sound{
		Id:       id,
		Name:     s.Name,
		FileName: s.FileName,
	}, nil
}
