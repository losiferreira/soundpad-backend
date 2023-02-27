package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"soundpad-backend/dals"
	authHandlers "soundpad-backend/flow/auth/handlers"
	"soundpad-backend/flow/auth/use_cases"
	soundHandlers "soundpad-backend/flow/sound/handlers"
	"soundpad-backend/providers"
)

func main() {
	_ = godotenv.Load(".env")

	ctx := context.Background()

	// providers
	bun := providers.NewBunDatabase().Setup()
	providers.NewGoth().Setup()

	// dals
	userDal := dals.NewUserDal(bun.Db, ctx)

	// useCases
	authUseCase := use_cases.NewAuthUsecase(userDal)

	// handlers
	googleHandler := authHandlers.NewGoogleHandler(authUseCase)
	soundHandler := soundHandlers.NewSoundHandler()
	rootHandler := authHandlers.NewRootHandler()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler.HandleRoot)
	//Auth
	r.HandleFunc("/auth/{provider}", googleHandler.HandleGoogleLogin)
	r.HandleFunc("/auth/{provider}/callback", googleHandler.HandleGoogleCallback)
	//Sound
	r.HandleFunc("/sound", soundHandler.HandlerUploadSound).Methods("POST")

	log.Println("Listening to port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Failed to listen and serve")
	}
}
