package main

import (
	lib "auth/lib"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("OAuth2-Token")
	if cookie == nil || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		w.Header().Add("Authorization", "serviceToken")
		w.WriteHeader(http.StatusOK)
	}
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {

	var sessionInstance = lib.SessionObjGen()

	var cookie http.Cookie = http.Cookie{
		Name:     sessionInstance.Cookie.Name,
		Value:    sessionInstance.Cookie.Value,
		Path:     sessionInstance.Cookie.Path,
		MaxAge:   sessionInstance.Cookie.MaxAge,
		HttpOnly: sessionInstance.Cookie.HttpOnly,
		Secure:   sessionInstance.Cookie.Secure,
		SameSite: sessionInstance.Cookie.SameSite,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/verify", verifyHandler)

	err := http.ListenAndServe(":4128", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
