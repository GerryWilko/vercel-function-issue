package handler

import "net/http"

func HandleUserAuthDIFFERENT(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
