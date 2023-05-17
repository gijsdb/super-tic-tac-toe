package controller

import (
	"net/http"
	"time"
)

func CreateCookie(name, value string, expiry time.Time, serverOnly bool) *http.Cookie {
	return &http.Cookie{Name: name, Value: value, Expires: expiry, SameSite: http.SameSiteStrictMode, Path: "/", Secure: true, HttpOnly: serverOnly}
}
