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

func main() {
	http.HandleFunc("/auth", authHandler)

	err := http.ListenAndServe(":4128", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
