package controllers

import (
	"net/http"
	"os"

	"github.com/MatthewZito/goldmund-sh-api/shared"
)

// Health is a liveness check that returns the server's current status
func Health(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		shared.FError(w, http.StatusBadRequest, "Healthcheck failed")
	}
	shared.FResponse(w, http.StatusOK, map[string]string{"server": name, "result": "success"})
}
