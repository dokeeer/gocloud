package auth

import (
	"encoding/json"
	"fmt"
	"gocloud/models"
	"gocloud/services/db"
	"net/http"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("gocloud_secret_key_179917")

func checkPassword(username, password string) bool {
	tmpDB, err := bolt.Open("../../storages/test.db", 0600, nil)
	if err != nil {
		return false
	}
	defer tmpDB.Close()
	users, err := db.GetStorageData(tmpDB, "personal")
	storedPassword, ok := users[username]
	if !ok {
		return false
	}
	return storedPassword == password
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if checkPassword(user.Username, user.Password) == false {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := createToken(user.Username)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
