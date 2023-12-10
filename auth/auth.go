package main

import (
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
		w.Header().Add("Authorization", "nfn61xt4udbi793p")
		w.WriteHeader(http.StatusOK)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "OAuth2-Token",
		Value:    "nfn61xt4udbi793p",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "http://localhost:3000/", 302)
}

func main() {
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe(":4128", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
