package main

import (
	"LinkShortening/build"
	"LinkShortening/config"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
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

	connectionPostgres, errDB := build.CreateDb(dbConfig)
	if errDB != nil {
		fmt.Println(errDB.Error())
		os.Exit(2)
	}
	defer connectionPostgres.Close()

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

	errStart := fasthttp.ListenAndServe(":5000", middlewareApi.LogURL(myRouter.Handler))
	if errStart != nil {
		fmt.Println(errStart.Error())
		os.Exit(2)
	}
}
