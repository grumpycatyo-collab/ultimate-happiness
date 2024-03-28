package testgrp

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler for dev
func (h Handlers) Test(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)

	statusCode := http.StatusOK
	h.Log.Infow("readiness",
		"statusCode", statusCode,
		"method", r.Method,
		"path", r.URL.Path,
		"remoteaddr", r.RemoteAddr)
}