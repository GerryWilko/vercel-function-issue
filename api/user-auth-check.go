package handler

import "net/http"

func HandleUserAuthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
