package providers

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"os"
)

type Goth struct {
}

func NewGoth() *Goth {
	return &Goth{}
}

func (g *Goth) Setup() {
	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = false

	gothic.Store = store

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_ID"),
			os.Getenv("GOOGLE_SECRET"),
			"http://localhost:8080/auth/google/callback",
			"email", "profile",
		),
	)
}
