package handlers

import (
	"net/http"
)

type SoundHandler struct {
}

func NewSoundHandler() *SoundHandler {
	return &SoundHandler{}
}

func (r *SoundHandler) HandlerUploadSound(w http.ResponseWriter, req *http.Request) {

}
