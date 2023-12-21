package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

const htmlElem = `
<h1>Authorized</h1>`

func testHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("OAuth2-Token")
	if cookie == nil || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		w.Header().Add("Authorization", "nfn61xt4udbi793p")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, htmlElem)
		return
	}
}

func main() {

	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":5603", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
