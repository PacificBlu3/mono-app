package sessionHandler

import {
	"encoding/json"
	"os"
	"time"
	"math/rand"
	"net/http"
	"errors"
}

var cookieStr string = cookieValGen(16)
var serviceToken string = cookieValGen(16)

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

function sessionObjGen() SessionObject {
	return SessionObject{
		Id:       1,
		Username: "test",
		Cookie:   sessionHandler.cookieObjGen(),
		LoggedAt: time.Now(),
	}
}

function cookieObjGen() CookieObject {
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

func cookieValGen(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
