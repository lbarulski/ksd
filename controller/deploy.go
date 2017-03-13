package controller

import (
	"fmt"
	"net/http"
	"os"
)

func Deploy(w http.ResponseWriter, r *http.Request) {
	if "POST" != r.Method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	token := r.URL.Query().Get("token")
	if !isTokenValid(token) {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// TODO: parse body & deploy!
	fmt.Println(r.Body)
}

func isTokenValid(token string) bool {
	validToken := os.Getenv("KSD_TOKEN")

	if len(validToken) < 1 {
		panic("ENV 'KSD_TOKEN' not set!")
	}

	return validToken == token
}
