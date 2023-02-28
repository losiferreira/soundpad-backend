package handlers

import (
	"log"
	"net/http"
	"soundpad-backend/providers"
)

func HandleValidateToken(
	w http.ResponseWriter,
	r *http.Request,
	jwt *providers.JwtHandler,
) (int64, error) {
	tokenString := r.Header.Get("token")
	token, err := jwt.GetToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Error getting token from header: %s", err)
		return -1, err
	}
	ownerId, err := jwt.GetUserIdFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Error parsing id from token: %s", err)
		return -1, err
	}
	return ownerId, nil
}
