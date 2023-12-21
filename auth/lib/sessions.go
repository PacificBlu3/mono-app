package lib

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var cookieStr string = CookieValGen(16)
var serviceToken string = CookieValGen(16)

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
