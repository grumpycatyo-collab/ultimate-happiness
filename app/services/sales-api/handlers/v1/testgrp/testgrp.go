package testgrp

import (
	"context"
	"github.com/grumpycatyo-collab/ultimate-happiness/foundation/web"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler for dev
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		//return web.NewShutdownError("restart service")
		panic("testing panic")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
