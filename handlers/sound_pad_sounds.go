package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"soundpad-backend/handlers/models"
	"soundpad-backend/use_cases"
	"strconv"
)

type SoundPadSoundsHandler struct {
	useCase *use_cases.SoundPadSoundsUseCase
}

func NewSoundPadSoundsHandler(
	useCase *use_cases.SoundPadSoundsUseCase,
) *SoundPadSoundsHandler {
	return &SoundPadSoundsHandler{
		useCase: useCase,
	}
}

func (s *SoundPadSoundsHandler) HandleCreateSoundPadSound(w http.ResponseWriter, r *http.Request) {
	soundPadSounds := &models.SoundPadToSound{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(soundPadSounds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding soundPadSounds: %s", err)
		return
	}

	err = s.useCase.CreateSoundPadSounds(soundPadSounds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error creating soundPadSound: %s", err)
		return
	}
}

func (s *SoundPadSoundsHandler) HandleRetrieveSoundPadSound(w http.ResponseWriter, r *http.Request) {
	stringSoudPadId := r.URL.Query().Get("soundPadId")
	soundPadId, err := strconv.ParseInt(stringSoudPadId, 10, 64)
	stringSoudId := r.URL.Query().Get("soundId")
	soundId, err := strconv.ParseInt(stringSoudId, 10, 64)

	if soundPadId <= 0 || soundId <= 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	soundPadSound := &models.SoundPadToSound{
		SoundPadId: json.Number(strconv.FormatInt(soundPadId, 10)),
		SoundId:    json.Number(strconv.FormatInt(soundId, 10)),
	}

	soundPadSound, err = s.useCase.RetrieveSoundPadSounds(soundPadSound)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(soundPadSound)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *SoundPadSoundsHandler) HandleDeleteSoundPadSound(w http.ResponseWriter, r *http.Request) {
	stringSoudPadId := r.URL.Query().Get("soundPadId")
	soundPadId, err := strconv.ParseInt(stringSoudPadId, 10, 64)
	stringSoudId := r.URL.Query().Get("soundId")
	soundId, err := strconv.ParseInt(stringSoudId, 10, 64)

	if soundPadId <= 0 || soundId <= 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error parsing params: %s", err)
		return
	}

	soundPadSound := &models.SoundPadToSound{
		SoundPadId: json.Number(strconv.FormatInt(soundPadId, 10)),
		SoundId:    json.Number(strconv.FormatInt(soundId, 10)),
	}

	err = s.useCase.DeleteSoundPadSounds(soundPadSound)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error deleting sound pad sound: %s", err)
		return
	}
}
