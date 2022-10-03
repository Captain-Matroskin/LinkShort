package application

import (
	"LinkShortening/internals/linkShort/orm"
)

type LinkShortAppInterface interface {
	CreateLinkShort(linkFull string) (string, error)
	TakeLinkShort(linkShort string) (string, error)
}

type LinkShortApp struct {
	Wrapper orm.LinkShortWrapperInterface
}

func (l *LinkShortApp) CreateLinkShort(linkFull string) (string, error) {

	return "shortLinkThis", l.Wrapper.CreateLinkShort(linkFull, "shortLinkThis") //TODO(N): refactor
}

func (l *LinkShortApp) TakeLinkShort(linkShort string) (string, error) {
	return l.Wrapper.TakeLinkShort(linkShort)
}
