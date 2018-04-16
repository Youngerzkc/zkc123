package jwt

import (
	"github.com/zkc123/dao"
	"errors"
	"github.com/zkc123/model"
	"fmt"
	// "net/http"
	// "time"
	jwt_lib "github.com/dgrijalva/jwt-go"
	// "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/zkc123/controller/common"
)
var (
	mysupersecretpassword = "mysecret"
)

// func  CreateToken() string {
// 	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
// 	// Set some claims
// 	token.Claims = jwt_lib.MapClaims{
// 		"id":  12,
// 		"exp": time.Now().Add(time.Hour * 1).Unix(),
// 	}
// 	// Sign and get the complete encoded token as a string
// 	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
// 	if err != nil {
// 		//c.JSON(500, gin.H{"message": "Could not generate token"})
// 		fmt.Println("token is error")
// 		return ""
// 	}
// 	//c.JSON(200, gin.H{"token": tokenString})
// 	return tokenString
	
// }

// func Auth(secret string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
// 			b := ([]byte(secret))
// 			return b, nil
// 		})

// 		if err != nil {
// 			c.AbortWithError(401, err)
// 		}
// 	}
// }
func getUser(c *gin.Context)(model.User,error)  {
	var user model.User
	tokenString,cookieErr:= c.Cookie("token")
	if cookieErr!=nil{
		return user,errors.New("未登录")
	}
	fmt.Println("get 用户的token ",tokenString)
	token,tokenErr:=jwt_lib.Parse(tokenString,func(token *jwt_lib.Token)(interface{},error){
		if _,ok:=token.Method.(*jwt_lib.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("Unexpected signed method:%v",token.Header["alg"])
		}
		return []byte(mysupersecretpassword),nil
	})
	fmt.Println("prase token",token.Valid)
	if tokenErr!=nil{
		return user,errors.New("未登录")
	}
	if claims,ok:=token.Claims.(jwt_lib.MapClaims);ok&&token.Valid{
		userID:=int(claims["id"].(float64))
		fmt.Println("UserId is",userID)
		var err error
		user,err:=dao.UserFromRedis(userID)
		if err!=nil{
			return user,errors.New("未登录")
		}
		return user,nil
	}
	return user,errors.New("未登录")
}
//必须是登录用户
func SigninRequired(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	var user model.User
	var err error
	if user, err = getUser(c); err != nil {
		SendErrJSON("未登录", c)
		return
	}
	c.Set("user", user)
	fmt.Println("必须是登录用户",user)
	c.Next()
}