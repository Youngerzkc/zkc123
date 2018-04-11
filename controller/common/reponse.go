package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func SendErrJSON(msg string,args ...interface{}){
	if len(args)==0{
		fmt.Println("缺少context")
	}
	var c *gin.Context
	//错误码定义 
	// var errNo=model.ErrorCode.ERROR
	if(len(args)==1){
		theCtx,ok:=args[0].(*gin.Context)
		if !ok{
			fmt.Println(msg,"需要msg和ctx")
			return 
		}
		c=theCtx
	}else if len(args)==2{
		theErrNo,ok:=args[0].(int)
		_=theErrNo
		if !ok{
			fmt.Println(msg,"需要msg,errNo和ctx")
			return 
		}
		// errNo=theErrNo
		theCtx,ok:=args[1].(*gin.Context)
		if !ok{
			fmt.Println(msg,"需要msg,errNo和ctx")
			return 
		}
		c=theCtx
	}
	c.JSON(401,gin.H{
		// "errNo":errNo,
		"msg":msg,
		"data":gin.H{},
	})
	
}

