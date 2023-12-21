package lib

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var cookieStr string = CookieValGen(16)
var serviceToken string = CookieValGen(16)

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

func SessionObjGen() SessionObject {
	return SessionObject{
		Id:       1,
		Username: "test",
		Cookie:   CookieObjGen(),
		LoggedAt: time.Now(),
	}
}

func CookieObjGen() CookieObject {
	return CookieObject{
		Name:     "OAuth2-Token",
		Value:    cookieStr,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}

func CookieValGen(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
