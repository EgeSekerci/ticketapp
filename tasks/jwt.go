package tasks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(user *User) (string, error) {
	// 	err := godotenv.Load()
	// 	shared.Check(err, "Error loading .env")

	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Minute * 15).Unix(),
		"userId":    user.Id,
		"userRole":  user.Role,
		"userName":  user.Name,
	}

	secret := os.Getenv("JWT_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Authorization")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		claims, err := validateJWT(cookie.Value)

		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		ctx := context.WithValue(r.Context(), claimsContextKey, claims)
		handlerFunc(w, r.WithContext(ctx))
	}
}

func validateJWT(tokenString string) (jwt.MapClaims, error) {
	// 	err := godotenv.Load()
	// 	shared.Check(err, "Error loading .env")

	secret := os.Getenv("JWT_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("Expired token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}
