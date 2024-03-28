package web

type Middleware func(Handler) Handler

func wrapMiddleware(mw []Middleware, handler Handler) Handler {

	// Loop backwards through  the middleware invoking each one
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}
	return handler
}
