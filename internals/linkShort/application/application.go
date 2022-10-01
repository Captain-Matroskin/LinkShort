package application

import (
	"LinkShortening/internals/linkShort/orm"
)

type LinkShortAppInterface interface {
	CreateLinkShort()
	TakeLinkShort()
}

type LinkShortApp struct {
	Wrapper orm.LinkShortWrapperInterface
}

func (l *LinkShortApp) CreateLinkShort() {

}

func (l *LinkShortApp) TakeLinkShort() {

}
