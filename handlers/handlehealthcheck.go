package handlers

import (
	"net/http"
)

/*
HandleHealthCheck ...
*/
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"alive"}`))
}
