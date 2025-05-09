package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"

	"ticketapp/tasks"
)

func TestLoginHandler(t *testing.T) {
	secretKey := "randomsecretkey"

	email := "test@test.com"
	password := "test123"

	form := url.Values{}

	form.Add("email", email)
	form.Add("password", password)

	req, err := http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(tasks.Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	if len(rr.Header().Get("Set-Cookie")) == 0 {
		t.Errorf("No cookie was set")
	}

	cookie := rr.Result().Cookies()[0]
	if cookie == nil {
		t.Fatal(err)
	}

	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		t.Fatal(err)
	}

	if !token.Valid {
		t.Errorf("Token is not valid")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["userId"] != float64(6) {
			t.Errorf("claims userId is incorrect")
		}
	} else {
		t.Errorf("Unable to parse token claims")
	}
}
