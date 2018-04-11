package jwt

import (
	"github.com/zkc123/dao"
	"errors"
	"github.com/zkc123/model"
	"fmt"
	"time"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)
var (
	mysupersecretpassword = "mysecret"
)

func  CreateToken() string {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"Id":  "Christopher",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
	if err != nil {
		//c.JSON(500, gin.H{"message": "Could not generate token"})
		fmt.Println("token is error")
		return ""
	}
	//c.JSON(200, gin.H{"token": tokenString})
	return tokenString
	
}

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}
func getUser(c *gin.Context)(model.User,error)  {
	var user model.User
	tokenString,cookieErr:=c.Cookie("token")
	if cookieErr!=nil{
		return user,errors.New("未登录")
	}
	token,tokenErr:=jwt_lib.Parse(tokenString,func(token *jwt_lib.Token)(interface{},error){
		if _,ok:=token.Method.(*jwt_lib.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("Unexpected signed method:%v",token.Header["alg"])
		}
		return []byte(mysupersecretpassword),nil
	})
	if tokenErr!=nil{
		return user,errors.New("未登录")
	}
	if claims,ok:=token.Claims.(jwt_lib.MapClaims);ok&&token.Valid{
		userID:=int(claims["id"].(float64))
		var err error
		user,err:=dao.UserFromRedis(userID)
		if err!=nil{
			return user,errors.New("未登录")
		}
		return user,nil
	}
	return user,errors.New("未登录")
}