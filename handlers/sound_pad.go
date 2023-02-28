package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"soundpad-backend/handlers/models"
	"soundpad-backend/providers"
	"soundpad-backend/use_cases"
	"strconv"
)

type SoundPadHandler struct {
	jwt     *providers.JwtHandler
	useCase *use_cases.SoundPadUseCase
}

func NewSoundPadHandler(
	jwt *providers.JwtHandler,
	useCase *use_cases.SoundPadUseCase,
) *SoundPadHandler {
	return &SoundPadHandler{
		jwt:     jwt,
		useCase: useCase,
	}
}

func (s *SoundPadHandler) HandleCreateSoundPad(w http.ResponseWriter, r *http.Request) {
	soundPad := &models.SoundPad{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(soundPad)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding soundPad: %s", err)
		return
	}

	ownerId, err := HandleValidateToken(w, r, s.jwt)
	if err != nil {
		return
	}

	err = s.useCase.CreateSoundPad(ownerId, soundPad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error uploading soundPad: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SoundPadHandler) HandleRetrieveSoundPad(w http.ResponseWriter, r *http.Request) {
	stringId := r.URL.Query().Get("id")
	intId, err := strconv.ParseInt(stringId, 10, 64)

	if intId < 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	soundPad, err := s.useCase.RetrieveSoundPad(intId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(soundPad)
	if err != nil {
		log.Printf("Error encoding sondpad: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SoundPadHandler) HandleUpdateSoundPad(w http.ResponseWriter, r *http.Request) {
	soundPad := &models.SoundPad{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(soundPad)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding soundPad: %s", err)
		return
	}

	soundPad, err = s.useCase.UpdateSoundPad(soundPad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error uploading soundPad_pad: %s", err)
		return
	}

	err = json.NewEncoder(w).Encode(soundPad)
	if err != nil {
		log.Printf("Error encoding sondpad: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SoundPadHandler) HandleDeleteSoundPad(w http.ResponseWriter, r *http.Request) {
	stringId := r.URL.Query().Get("id")
	intId, err := strconv.ParseInt(stringId, 10, 64)

	if intId < 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.useCase.DeleteSoundPad(intId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
