package main

import (
	"LinkShortening/build"
	"LinkShortening/config"
	proto "LinkShortening/internals/proto"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	runServer()
}

func runServer() {
	errConf, configRes := build.InitConfig()
	if errConf != nil {
		println(errConf.Error())
		os.Exit(1)
	}
	configMain := configRes[0].(config.MainConfig)
	configDB := configRes[1].(config.DBConfig)

	connectionPostgres, errDB := build.CreateConn(configDB.Db)
	if errDB != nil {
		fmt.Println(errDB.Error())
		os.Exit(2)
	}
	defer connectionPostgres.Close()

	errCreateDB := build.CreateDB(connectionPostgres)
	if errCreateDB != nil {
		fmt.Println(errDB.Error())
		os.Exit(2)
	}

	startStructure := build.SetUp(connectionPostgres)

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
		println(errListen.Error())
		os.Exit(1)
	}
	server := grpc.NewServer()

	proto.RegisterLinkShortServiceServer(server, &startStructure.LinkShortManager)

	go func() {
		errServ := server.Serve(listen)
		if errServ != nil {
			println(errServ.Error())
			os.Exit(1)
		}
	}()

	addresHttp := ":" + configMain.Main.PortHttp

	errStart := fasthttp.ListenAndServe(addresHttp, middlewareApi.LogURL(myRouter.Handler))
	if errStart != nil {
		fmt.Println(errStart.Error())
		os.Exit(2)
	}

}
