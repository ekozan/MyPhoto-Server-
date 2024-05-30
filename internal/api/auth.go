package api

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// @Summary Login
// @Description User login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body Credentials true "Login credentials"
// @Success 200 {string} string "Token"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if creds.Username != "user" || creds.Password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		UserID:   "user_id", // Replace with the actual user ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey := []byte(viper.GetString("jwt.secret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
