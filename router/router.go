package router

import (
	"github.com/zkc123/controller/common"
	"github.com/zkc123/controller/user"
	"github.com/zkc123/middler/jwt"
	// "github.com/zkc123/middler/logs"
	"github.com/zkc123/controller/test"
	"fmt"
	"github.com/gin-gonic/gin"

)
func Route(router *gin.Engine) {
	fmt.Println("start router...")
	router.GET("/test/txt",test.Test)
	
}
func RouteJwt(route *gin.Engine)  {

	fmt.Println("start jwt router...")
	route.GET("/jwt",test.TestJwt)
}
func RouterUser(router *gin.Engine){
	fmt.Println("user user router")
	r:=router.Group("/user",jwt.RefreshTokenCookie)
	{
		r.POST("/signin",user.Signin)//登录
		r.POST("/signup",user.Signup)//注册
		r.GET("/",test.Test)
		r.POST("/upload",common.UploadHandler)
	}
}