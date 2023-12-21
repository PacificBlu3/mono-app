package lib

import (
	"net/http"
	"time"
)

type CookieObject struct {
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	Path     string        `json:"path"`
	MaxAge   int           `json:"maxAge"`
	HttpOnly bool          `json:"httpOnly"`
	Secure   bool          `json:"secure"`
	SameSite http.SameSite `json:"sameSite"`
}

type SessionObject struct {
	Id       int          `json:"id"`
	Username string       `json:"username"`
	Cookie   CookieObject `json:"cookie"`
	LoggedAt time.Time    `json:"LoginTime"`
}
