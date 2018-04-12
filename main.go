package main

import (
	"github.com/zkc123/middler/jwt"
	"io"
	"github.com/zkc123/middler/logs"
	"log"
	"os/signal"
	"os"
	"fmt"
	"github.com/zkc123/router"
	"github.com/gin-gonic/gin"
)

func main() {
    // gin.SetMode("release")
	var outLog io.Writer
	var fileLog string="/tmp/zkc123.log"
	outLog=logs.FileLogs(fileLog)
	ginMiddleHandLogs:=logs.NewWithWriter(outLog)
	r:=gin.Default()
	r.Use(ginMiddleHandLogs)
	router.Route(r)
	fmt.Println("token is ",jwt.CreateToken())
	// r.Use(jwt.Auth("mysecret"))
	// r.Use(gin.Recovery())
	// router.RouteJwt(r)
	router.RouterUser(r)
	go func()  {
		r.Run(":8080")
	}()
	//优雅的关闭服务
	quit:=make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("shutdown server.... ")
}