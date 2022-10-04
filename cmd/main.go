package main

import (
	"LinkShortening/build"
	"LinkShortening/config"
	errPkg "LinkShortening/internals/myerror"
	proto "LinkShortening/internals/proto"
	"LinkShortening/internals/util"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	runServer()
}

func runServer() {
	var logger util.Logger
	logger.Log = util.NewLogger("./logs.txt")

	defer func(loggerErrWarn errPkg.MultiLoggerInterface) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(2)
		}
	}(logger.Log)

	errConfig, configRes := build.InitConfig()
	if errConfig != nil {
		logger.Log.Errorf("%s", errConfig.Error())
		os.Exit(1)
	}
	configMain := configRes[0].(config.MainConfig)
	configDB := configRes[1].(config.DBConfig)

	connectionPostgres, errDB := build.CreateConn(configDB.Db)
	if errDB != nil {
		logger.Log.Errorf("Err connect database: %s", errDB.Error())
		os.Exit(2)
	}
	defer connectionPostgres.Close()

	errCreateDB := build.CreateDB(connectionPostgres)
	if errCreateDB != nil {
		logger.Log.Errorf("err create database: %s", errCreateDB.Error())
		os.Exit(2)
	}

	startStructure := build.SetUp(connectionPostgres, logger.Log)

	linkShortApi := startStructure.LinkShort
	middlewareApi := startStructure.Middle

	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	linkShort := versionGroup.Group("/linkShort")

	linkShort.POST("/", linkShortApi.CreateLinkShortHandler)
	linkShort.GET("/", linkShortApi.TakeLinkShortHandler)
	//myRouter.GET("/health", )
	addresGrpc := configMain.Main.HostGrpc + ":" + configMain.Main.PortGrpc

	listen, errListen := net.Listen(configMain.Main.Network, addresGrpc)
	if errListen != nil {
		logger.Log.Errorf("Server listen grpc error: %v", errListen)
		os.Exit(1)
	}
	server := grpc.NewServer()

	proto.RegisterLinkShortServiceServer(server, &startStructure.LinkShortManager)

	go func() {
		logger.Log.Infof("Listen in %s", addresGrpc)
		errServ := server.Serve(listen)
		if errServ != nil {
			logger.Log.Errorf("Server serv grpc error: %v", errServ)
			os.Exit(1)
		}

	}()

	addresHttp := ":" + configMain.Main.PortHttp

	logger.Log.Infof("Listen in 127:0.0.1%s", addresHttp)
	errStart := fasthttp.ListenAndServe(addresHttp, middlewareApi.LogURL(myRouter.Handler))
	if errStart != nil {
		logger.Log.Errorf("Listen and server http error: %v", errStart)
		os.Exit(2)
	}

}
