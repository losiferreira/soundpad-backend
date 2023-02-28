package models

import (
	"soundpad-backend/dals/entity"
	"strconv"
)

type Sound struct {
	Id       string
	Name     string
	FileName string
}

func (s *Sound) FromEntity(sound *entity.Sound) *Sound {
	s.Id = strconv.FormatInt(sound.Id, 10)
	s.Name = sound.FileName
	s.FileName = sound.FileName
	return s
}

func (s *Sound) ToEntity() (*entity.Sound, error) {
	id, err := strconv.ParseInt(s.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	return &entity.Sound{
		Id:       id,
		Name:     s.Name,
		FileName: s.FileName,
	}, nil
}
