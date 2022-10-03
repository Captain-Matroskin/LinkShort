package build

import (
	apiLink "LinkShortening/internals/linkShort/api"
	"LinkShortening/internals/linkShort/application"
	"LinkShortening/internals/linkShort/orm"
	apiMiddle "LinkShortening/internals/middleware/api"
)

type installSetUp struct {
	LinkShort apiLink.LinkShortApi
	Middle    apiMiddle.MiddlewareApi
}

func SetUp(connectionDB orm.ConnectionInterface) *installSetUp {
	linkShortWrapper := orm.LinkShortWrapper{
		Conn: connectionDB,
	}
	linkShortApp := application.LinkShortApp{
		Wrapper: &linkShortWrapper,
	}
	linkShortApi := apiLink.LinkShortApi{
		Application: &linkShortApp,
	}
	var _ apiLink.LinkShortApiInterface = &linkShortApi

	middlewareApi := apiMiddle.MiddlewareApi{}
	var _ apiMiddle.MiddlewareApiInterface = &middlewareApi

	var result installSetUp
	result.LinkShort = linkShortApi

	return &result
}
