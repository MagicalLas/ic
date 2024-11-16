package api

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"gomod.usaken.org/ic/monitor"
	"time"
)

type FastHTTPApiHandler struct {
	serverName string
	logger     monitor.Logger
}

func (h *FastHTTPApiHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// TODO: fiber로 구현체를 변경하는게 더 좋은 선택인걸 알지만,
	//  fiber가 갖고있는 overhead를 가져오고싶지 않았기에
	//  일부 기능을 직접 대충 구현합니다.
	//  추후에는 fiber로 옮길 예정입니다.
	//h.logger.Info("request on")
	monitor.CollectHTTPRequest("/:all")

	path := ctx.Path()
	if bytes.HasPrefix(path, contentsPrefix) {
		ls := ctx.Request.URI().LastPathSegment()
		if bytes.HasPrefix(ls, impressionPrefix) {
			h.IncreaseImpression(ctx)
			monitor.CollectHTTPResponse("/IncreaseImpression", 200, time.Since(ctx.Time()))
			return
		}
		if bytes.HasPrefix(ls, clickPrefix) {
			h.IncreaseClick(ctx)
			monitor.CollectHTTPResponse("/IncreaseClick", 200, time.Since(ctx.Time()))
			return
		}
		return
	}
}

// hot path에서의 메모리 할당을 피하기 위하여 미리 선언.
var contentsPrefix = []byte("/contents")
var impressionPrefix = []byte("imp")
var clickPrefix = []byte("click")

func (h *FastHTTPApiHandler) IncreaseImpression(ctx *fasthttp.RequestCtx) {
	monitor.CollectHTTPRequest("/IncreaseImpression")
	//h.logger.Info("increase imp\n")
}

func (h *FastHTTPApiHandler) IncreaseClick(ctx *fasthttp.RequestCtx) {
	monitor.CollectHTTPRequest("/IncreaseClick")
	//h.logger.Info("increase click\n")
}

func (h *FastHTTPApiHandler) GetContents(ctx *fasthttp.RequestCtx) {

}
