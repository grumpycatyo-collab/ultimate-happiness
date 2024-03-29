package mid

import (
	"context"
	"github.com/grumpycatyo-collab/ultimate-happiness/foundation/web"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// using closures
func Logger(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v, err := web.GetValues(ctx)
			if err != nil {
				return err // web.NewShutdownError("web value missing")
			}
			log.Infow("request started", "traceid", v.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			err = handler(ctx, w, r)

			log.Infow("request completed", "traceid", v.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))
			return err
		}
		return h
	}
	return m
}
