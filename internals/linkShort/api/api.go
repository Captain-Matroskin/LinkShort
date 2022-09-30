package api

import "github.com/valyala/fasthttp"

type LinkShortApiInterface interface {
	CreateLinkShortHandler(ctx *fasthttp.RequestCtx)
	TakeLinkShortHandler(ctx *fasthttp.RequestCtx)
}

type LinkShortApi struct {
}

func (l *LinkShortApi) CreateLinkShortHandler(ctx *fasthttp.RequestCtx) {

}

func (l *LinkShortApi) TakeLinkShortHandler(ctx *fasthttp.RequestCtx) {

}

