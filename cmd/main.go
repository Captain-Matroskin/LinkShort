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
	dbConfig := config.Database{ //TODO(N): refactor in config
		Host:       "127.0.0.1",
		Port:       "5432",
		UserName:   "Captain-matroskin",
		Password:   "74tbr6r54f78",
		SchemaName: "postgres",
	}

	connectionPostgres, errDB := build.CreateConn(dbConfig)
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

	listen, errListen := net.Listen("tcp", "127.0.0.1:8081")
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

	errStart := fasthttp.ListenAndServe(":5000", middlewareApi.LogURL(myRouter.Handler))
	if errStart != nil {
		fmt.Println(errStart.Error())
		os.Exit(2)
	}

}
