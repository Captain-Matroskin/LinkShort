//go:generate mockgen -destination=mocks/api.go -package=mocks LinkShortening/internals/myerror MultiLoggerInterface
package api

import (
	errPgk "LinkShortening/internals/myerror"
	"github.com/valyala/fasthttp"
	"math"
)

type MiddlewareApiInterface interface {
	LogURL(h fasthttp.RequestHandler) fasthttp.RequestHandler
}

type MiddlewareApi struct {
	ReqId  int
	Logger errPgk.MultiLoggerInterface
}

func (m *MiddlewareApi) LogURL(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		if m.ReqId == math.MaxInt {
			m.ReqId = 0
		}
		m.ReqId++
		m.Logger.Infof("Method: %s, URL: %s, requestId: %d", string(ctx.Method()), ctx.URI(), m.ReqId)
		ctx.SetUserValue("reqId", m.ReqId)

		h(ctx)
	})
}
