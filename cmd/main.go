package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	runServer()
}

func runServer() {
	myRouter := router.New()
	apiGroup := myRouter.Group("/api")
	versionGroup := apiGroup.Group("/v1")
	linkShort := versionGroup.Group("/linkShort")
	_ = linkShort

	//linkShort.POST("/", )
	//linkShort.GET("/", )
	//myRouter.GET("/health", _)

	errStart := fasthttp.ListenAndServe(":5000", myRouter.Handler)
	if errStart != nil {
		fmt.Println(errStart.Error())
		os.Exit(2)
	}
}
