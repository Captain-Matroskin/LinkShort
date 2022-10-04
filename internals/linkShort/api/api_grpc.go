package api

import (
	"LinkShortening/internals/linkShort/application"
	errPkg "LinkShortening/internals/myerror"
	proto "LinkShortening/internals/proto"
	"context"
	"net/http"
)

type LinkShortManagerInterface interface {
	CreateLinkShort(ctx context.Context, linkFullIn *proto.LinkFull) (*proto.ResultLinkShort, error)
	TakeLinkFull(ctx context.Context, linkShortIn *proto.LinkShort) (*proto.ResultLinkFull, error)
}

type LinkShortManager struct {
	Application application.LinkShortAppInterface
}

func (l *LinkShortManager) CreateLinkShort(ctx context.Context, linkFullIn *proto.LinkFull) (*proto.ResultLinkShort, error) {
	checkError := &errPkg.CheckError{}

	linkShortOut, errIn := l.Application.CreateLinkShortApp(linkFullIn.LinkFull)

	errOut, resultOut, codeHTTP := checkError.CheckErrorCreateLinkShortGrpc(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrCheck:
			return &proto.ResultLinkShort{Error: resultOut, StatusCode: int64(codeHTTP)}, nil
		case errPkg.ErrInternal:
			return &proto.ResultLinkShort{}, &errPkg.MyErrors{Text: resultOut}

		}
	}

	return &proto.ResultLinkShort{StatusCode: http.StatusOK, Body: &proto.LinkShort{LinkShort: linkShortOut}}, nil

}

func (l *LinkShortManager) TakeLinkFull(ctx context.Context, linkShortIn *proto.LinkShort) (*proto.ResultLinkFull, error) {
	checkError := &errPkg.CheckError{}

	linkFullOut, errIn := l.Application.TakeLinkFullApp(linkShortIn.LinkShort)

	errOut, resultOut, codeHTTP := checkError.CheckErrorCreateLinkShortGrpc(errIn)
	if errOut != nil {
		switch errOut.Error() {
		case errPkg.ErrCheck:
			return &proto.ResultLinkFull{Error: resultOut, StatusCode: int64(codeHTTP)}, nil
		case errPkg.ErrInternal:
			return &proto.ResultLinkFull{}, &errPkg.MyErrors{Text: resultOut}

		}
	}

	return &proto.ResultLinkFull{StatusCode: http.StatusOK, Body: &proto.LinkFull{LinkFull: linkFullOut}}, nil
}
