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

// IsValidServiceType checks if the provided service type is valid

func IsValidServiceType(serviceType string) bool {

	validTypes := []string{"home cleaning", "cleaning home", "clean at home", "clean at store", "store cleaning", "cleaning store", "Home Cleaning"}

	serviceTypeLower := strings.ToLower(serviceType)
	for _, t := range validTypes {
		if strings.ToLower(t) == serviceTypeLower {
			return true
		}
	}

	return false

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
