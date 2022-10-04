package build

import (
	"LinkShortening/config"
	api "LinkShortening/internals/linkShort/api"
	"LinkShortening/internals/linkShort/application"
	"LinkShortening/internals/linkShort/orm"
	apiMiddle "LinkShortening/internals/middleware/api"
	errPkg "LinkShortening/internals/myerror"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfType     = "yml"
	ConfPath     = "./config/"
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

func InitConfig() (error, []interface{}) {
	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameMain)
	errRead := viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.MyErrors{
			Text: errRead.Error(),
		}, nil
	}
	mainConfig := config.MainConfig{}
	errUnmarshal := viper.Unmarshal(&mainConfig)
	if errUnmarshal != nil {
		return &errPkg.MyErrors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.MyErrors{
			Text: errRead.Error(),
		}, nil
	}
	dbConfig := config.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.MyErrors{
			Text: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, mainConfig)
	result = append(result, dbConfig)

	return nil, result
}
