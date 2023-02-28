package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"soundpad-backend/dals"
	"soundpad-backend/dals/entity"
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
	aws := providers.NewAws().Setup()
	jwt := providers.NewJwtHandler()

	// entity setup
	entity.NewEntityRegistration(bun.Db).Setup()

	// dals
	userDal := dals.NewUserDal(bun.Db, ctx)
	soundDal := dals.NewSoundDal(bun.Db, ctx)
	soundPadDal := dals.NewSoundPadDal(bun.Db, ctx)
	soundPadSoundsDal := dals.NewSoundPadSoundsDal(bun.Db, ctx)

	// useCases
	authUseCase := use_cases.NewAuthUsecase(jwt, userDal)
	soundUseCase := use_cases.NewSoundUseCase(aws, soundDal)
	soundPadUseCase := use_cases.NewSoundPadUseCase(soundPadDal)
	soundPadSoundsUseCase := use_cases.NewSoundPadSoundsUseCase(soundPadSoundsDal)

	// handlers
	googleHandler := handlers.NewGoogleHandler(authUseCase)
	soundHandler := handlers.NewSoundHandler(soundUseCase)
	soundPadHandler := handlers.NewSoundPadHandler(jwt, soundPadUseCase)
	soundPadSoundsHandler := handlers.NewSoundPadSoundsHandler(soundPadSoundsUseCase)
	rootHandler := handlers.NewRootHandler()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler.HandleRoot)
	//Auth
	r.HandleFunc("/auth/{provider}", googleHandler.HandleGoogleLogin)
	r.HandleFunc("/auth/{provider}/callback", googleHandler.HandleGoogleCallback)
	//Sound
	r.HandleFunc("/sounds", soundHandler.HandleCreateSound).Methods("POST")
	r.HandleFunc("/sounds", soundHandler.HandleRetrieveSound).Methods("GET")
	r.HandleFunc("/sounds", soundHandler.HandleDeleteSound).Methods("DELETE")
	//SoundPad
	r.HandleFunc("/soundPads", soundPadHandler.HandleCreateSoundPad).Methods("POST")
	r.HandleFunc("/soundPads", soundPadHandler.HandleRetrieveSoundPad).Methods("GET")
	r.HandleFunc("/soundPads", soundPadHandler.HandleUpdateSoundPad).Methods("PUT")
	r.HandleFunc("/soundPads", soundPadHandler.HandleDeleteSoundPad).Methods("DELETE")
	//SoundPadSounds
	r.HandleFunc("/soundPads/sounds", soundPadSoundsHandler.HandleCreateSoundPadSound).Methods("POST")
	r.HandleFunc("/soundPads/sounds", soundPadSoundsHandler.HandleRetrieveSoundPadSound).Methods("GET")
	r.HandleFunc("/soundPads/sounds", soundPadSoundsHandler.HandleDeleteSoundPadSound).Methods("DELETE")

	log.Println("Listening to port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Failed to listen and serve")
	}
}
