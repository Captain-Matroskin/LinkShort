package api

import (
	"github.com/valyala/fasthttp"
	"math"
)

type MiddlewareApiInterface interface {
	LogURL(h fasthttp.RequestHandler) fasthttp.RequestHandler
}

type MiddlewareApi struct {
	ReqId int
}

func (m *MiddlewareApi) LogURL(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		if m.ReqId == math.MaxInt {
			m.ReqId = 0
		}
		m.ReqId++
		//TODO(M): logger
		ctx.SetUserValue("reqId", m.ReqId)

		h(ctx)
	})
}
