package utils

import (
	"net/http"
	"time"
)

func CreateCookie(name string, value string, expiresIn time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  expiresIn,
		SameSite: http.SameSiteStrictMode,
	}
}

func DeleteCookie(name string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
}
