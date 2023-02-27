package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"soundpad-backend/dals"
	"soundpad-backend/handlers"
	"soundpad-backend/providers"
	"soundpad-backend/use_cases"
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
	googleHandler := handlers.NewGoogleHandler(authUseCase)
	rootHandler := handlers.NewRootHandler()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler.HandleRoot)
	r.HandleFunc("/auth/{provider}", googleHandler.HandleGoogleLogin)
	r.HandleFunc("/auth/{provider}/callback", googleHandler.HandleGoogleCallback)
	log.Println("Listening to port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Failed to listen and serve")
	}
}
