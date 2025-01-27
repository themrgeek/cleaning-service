package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(nil, w, http.StatusInternalServerError, "Failed to process payload")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(c *gin.Context, w http.ResponseWriter, code int, message string) {
	if c != nil {
		c.JSON(code, gin.H{"error": message})
		return
	}
	RespondWithJSON(w, code, map[string]string{"error": message})
}
