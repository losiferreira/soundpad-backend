package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"soundpad-backend/use_cases"
	"strconv"
)

type SoundHandler struct {
	useCase *use_cases.SoundUseCase
}

func NewSoundHandler(
	useCase *use_cases.SoundUseCase,
) *SoundHandler {
	return &SoundHandler{
		useCase: useCase,
	}
}

func (s *SoundHandler) HandleCreateSound(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4 << 20) // 32 MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("File must be 4MB or smaller")
		return
	}

	// Get the MP3 file from the form data
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error forming sound_pad: %s", err)
		return
	}
	defer file.Close()

	name := r.Form.Get("name")

	err = s.useCase.CreateSound(file, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error uploading sound_pad: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SoundHandler) HandleRetrieveSound(w http.ResponseWriter, r *http.Request) {
	stringId := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(stringId)

	if intId < 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sound, err := s.useCase.RetrieveSound(intId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(sound)
	w.WriteHeader(http.StatusOK)
}

func (s *SoundHandler) HandleDeleteSound(w http.ResponseWriter, r *http.Request) {
	stringId := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(stringId)

	if intId < 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.useCase.DeleteSound(intId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
