package shared

import (
	"encoding/json"
	"net/http"
)

func FError(w http.ResponseWriter, code int, msg string) {
	FResponse(w, code, map[string]string{"error": msg})
}

func FResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Clacks-Overhead", "RnJlZVNwZWVjaAo=")
	w.Header().Set("X-Powered-By", "goldmund.sh/2.0")
	w.WriteHeader(code)
	w.Write(response)
}
