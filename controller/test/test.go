package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	
)
func Test (r *gin.Context)  {		
	fmt.Println("test")
	r.JSON(200,"hello world")
}
func TestJwt(r *gin.Context )  {
	
	fmt.Println("testJwt")
	r.JSON(200,"hello testJwt")
}