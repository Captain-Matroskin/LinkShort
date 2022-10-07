//go:generate mockgen -destination=mocks/application.go -package=mocks LinkShortening/internals/linkShort/orm LinkShortWrapperInterface
package application

import (
	"LinkShortening/internals/linkShort/orm"
	errPkg "LinkShortening/internals/myerror"
	"LinkShortening/internals/util"
	"github.com/aidarkhanov/nanoid"
)

type LinkShortAppInterface interface {
	CreateLinkShortApp(linkFull string) (string, error)
	TakeLinkFullApp(linkShort string) (string, error)
}

type LinkShortApp struct {
	Wrapper orm.LinkShortWrapperInterface
}

func (l *LinkShortApp) CreateLinkShortApp(linkFull string) (string, error) {
	generateLinkShort, err := nanoid.Generate(nanoid.DefaultAlphabet, util.LenLinkShort)
	if err != nil {
		return "", &errPkg.MyErrors{
			Text: errPkg.LSHCreateLinkShortAppNotGenerate,
		}
	}
	generateLinkShort = util.LinkDomain + "/" + generateLinkShort

	return generateLinkShort, l.Wrapper.CreateLinkShort(linkFull, generateLinkShort)
}

func (l *LinkShortApp) TakeLinkFullApp(linkShort string) (string, error) {
	return l.Wrapper.TakeLinkFull(linkShort)
}
