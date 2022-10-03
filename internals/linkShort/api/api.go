package api

import (
	"LinkShortening/internals/linkShort"
	"LinkShortening/internals/linkShort/application"
	errPkg "LinkShortening/internals/myerror"
	"LinkShortening/internals/util"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

type LinkShortApiInterface interface {
	CreateLinkShortHandler(ctx *fasthttp.RequestCtx)
	TakeLinkShortHandler(ctx *fasthttp.RequestCtx)
}

type LinkShortApi struct {
	Application application.LinkShortAppInterface
}

func (l *LinkShortApi) CreateLinkShortHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		fmt.Println(errConvert.Error())
	}

	checkError := &errPkg.CheckError{
		RequestId: reqId,
	}

	var linkFullIn linkShort.LinkFull
	errUnmarshal := json.Unmarshal(ctx.Request.Body(), &linkFullIn)
	if errUnmarshal != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		fmt.Println(errUnmarshal.Error())
		return
	}

	linkShortOut, errIn := l.Application.CreateLinkShortApp(linkFullIn.Link)

	errOut, resultOut, codeHTTP := checkError.CheckErrorCreateLinkShort(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	request, errRequest := json.Marshal(&util.Result{
		Status: http.StatusCreated,
		Body: linkShort.ResponseLinkShort{
			LinkShort: linkShort.LinkShort{
				Link: linkShortOut,
			},
		},
	})
	if errRequest != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		println(errRequest.Error())
		return
	}

	ctx.Response.SetBody(request)
	json.NewEncoder(ctx)
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (l *LinkShortApi) TakeLinkShortHandler(ctx *fasthttp.RequestCtx) {
	reqIdCtx := ctx.UserValue("reqId")
	reqId, errConvert := util.InterfaceConvertInt(reqIdCtx)
	if errConvert != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errConvert.Error()))
		return
	}

	checkError := &errPkg.CheckError{
		RequestId: reqId,
	}

	var linkShortIn linkShort.LinkShort
	errUnmarshal := json.Unmarshal(ctx.Request.Body(), &linkShortIn)
	if errUnmarshal != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrUnmarshal))
		fmt.Println(errUnmarshal.Error())
		return
	}

	linkFullOut, errIn := l.Application.TakeLinkShortApp(linkShortIn.Link)

	errOut, resultOut, codeHTTP := checkError.CheckErrorTakeLinkShort(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrMarshal:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody([]byte(errPkg.ErrMarshal))
			return
		case errPkg.ErrCheck:
			ctx.Response.SetStatusCode(codeHTTP)
			ctx.Response.SetBody(resultOut)
			return
		}
	}

	request, errRequest := json.Marshal(&util.Result{
		Status: http.StatusCreated,
		Body: linkShort.ResponseLinkFull{
			LinkShort: linkShort.LinkFull{
				Link: linkFullOut,
			},
		},
	})
	if errRequest != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBody([]byte(errPkg.ErrEncode))
		println(errRequest.Error())
		return
	}

	ctx.Response.SetBody(request)
	json.NewEncoder(ctx)
	ctx.Response.SetStatusCode(http.StatusOK)

}
