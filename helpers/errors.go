package helpers

import (
	"encoding/json"
	"net/http"
)

// ResponseWithJSON control response with json format
func ResponseWithJSON(w http.ResponseWriter, status int, data interface{}) {
	result, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(result)
}
