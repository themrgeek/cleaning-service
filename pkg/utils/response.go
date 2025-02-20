package utils

import (
	"encoding/json"
	"html"
	"net/http"
	"strings"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to process payload")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SanitizeString removes potentially harmful characters and HTML escapes the input

func SanitizeString(input string) string {

	// Trim spaces

	sanitized := strings.TrimSpace(input)

	// Escape HTML special characters

	return html.EscapeString(sanitized)

}

func RespondWithError(w http.ResponseWriter, code int, message string) {

	RespondWithJSON(w, code, map[string]string{"error": message})
}
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(dst)

	if err != nil {

		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return err

	}

	return nil

}
