package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, value any, status int) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func WriteErrors(w http.ResponseWriter, errors map[string]string, status int) {
	WriteJSON(w, errors, status)
}

func WriteError(w http.ResponseWriter, err error, status int) {
	WriteErrors(w, map[string]string{"error": err.Error()}, status)
}

func WriteMsg(w http.ResponseWriter, msg string, status int) {
	WriteErrors(w, map[string]string{"msg": msg}, status)
}
