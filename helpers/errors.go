package helpers

import (
	"encoding/json"
	"net/http"
)

// ResponseWithErrs control pkgErrors handling for response
func ResponseWithErrs(w http.ResponseWriter, status int, msg string) {
	ResponseWithJSON(w, status, map[string]interface{}{"error": msg, "status": status})
}

// ResponseWithJSON control response with json format
func ResponseWithJSON(w http.ResponseWriter, status int, data interface{}) {
	result, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(result)
}
