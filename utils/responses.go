package utils

import (
	"encoding/json"
	"net/http"
)

// SendErr send a error message with status as response
func SendErr(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, map[string]string{"error": msg})
}

// SendJSON return a parsed json with status as response
func SendJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
