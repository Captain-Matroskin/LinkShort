package build

import (
	api "LinkShortening/internals/linkShort/api"
	"LinkShortening/internals/linkShort/application"
	"LinkShortening/internals/linkShort/orm"
	apiMiddle "LinkShortening/internals/middleware/api"
)

type installSetUp struct {
	LinkShort        api.LinkShortApi
	Middle           apiMiddle.MiddlewareApi
	LinkShortManager api.LinkShortManager
}

func SetUp(connectionDB orm.ConnectionInterface) *installSetUp {
	linkShortWrapper := orm.LinkShortWrapper{
		Conn: connectionDB,
	}
	linkShortApp := application.LinkShortApp{
		Wrapper: &linkShortWrapper,
	}
	linkShortApi := api.LinkShortApi{
		Application: &linkShortApp,
	}
	var _ api.LinkShortApiInterface = &linkShortApi

	middlewareApi := apiMiddle.MiddlewareApi{}
	var _ apiMiddle.MiddlewareApiInterface = &middlewareApi

	linkShortManager := api.LinkShortManager{
		Application: &linkShortApp,
	}
	var _ api.LinkShortManagerInterface = &linkShortManager

	var result installSetUp
	result.LinkShort = linkShortApi
	result.Middle = middlewareApi
	result.LinkShortManager = linkShortManager

	return &result
}
