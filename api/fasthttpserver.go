package api

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type FastHTTPApiHandler struct {
	foobar string
}

// request handler in net/api style, i.e. method bound to MyHandler struct.
func (h *FastHTTPApiHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func runServer() {
	myHandler := &FastHTTPApiHandler{
		foobar: "foobar",
	}
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
}
