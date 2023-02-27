package handlers

import (
	"encoding/json"
	"github.com/markbates/goth/gothic"
	"html/template"
	"log"
	"net/http"
	"soundpad-backend/flow/auth/handlers/models"
	"soundpad-backend/flow/auth/shared"
	"soundpad-backend/flow/auth/use_cases"
)

type GoogleHandler struct {
	authUseCase *use_cases.AuthUseCase
}

func NewGoogleHandler(
	authUseCase *use_cases.AuthUseCase,
) *GoogleHandler {
	return &GoogleHandler{
		authUseCase: authUseCase,
	}
}

func (g *GoogleHandler) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (g *GoogleHandler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Fatal(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Google user email: %s\n", user.Email)
	isTemplateTest := shared.GetOsBoolEnv("TEMPLATE_TEST")

	token, err := g.authUseCase.SignInOrSignUp(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isTemplateTest {
		t, _ := template.ParseFiles("test_templates/google_sign_success.html")
		_ = t.Execute(w, user)
	} else {
		result := models.AuthResponse{
			Token: token,
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
