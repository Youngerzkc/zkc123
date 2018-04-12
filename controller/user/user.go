package user

import (
	"time"
	"net/http"
	"github.com/zkc123/dao"
	"github.com/zkc123/model"
	"github.com/gin-gonic/gin/binding"
	"github.com/zkc123/controller/common"
	"fmt"
	"github.com/gin-gonic/gin"
	jwt_lib "github.com/dgrijalva/jwt-go"
)
//用户登录
func Signin(c *gin.Context){
	SendErrJSON:=common.SendErrJSON
	type UserNameLogin struct{
		SigninInput string `json:"signinInput" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}

	 var userNameLogin UserNameLogin
	 var signinInput string
	 var password string
	 fmt.Println("parmars is ",c.Query("loginType"))
	if c.Query("loginType")=="username"{
		if err:=c.ShouldBindWith(&userNameLogin,binding.JSON);err!=nil{
			// fmt.Println("JSON ",binding.JSON)
			fmt.Println(err.Error())
			fmt.Println("name iss isi", userNameLogin.SigninInput)
			fmt.Println("格式化参数")
			SendErrJSON("用户名或密码错误",c)
			return 
		}
		signinInput=userNameLogin.SigninInput
		password=userNameLogin.Password
	}
	var user model.User
	if err:=dao.SelectUserMysql(signinInput);err!=nil{
		SendErrJSON("帐号不存在",c)
		return 
	}
	userDao := &dao.User{}
	userDao.Pass=password
	userDao.Name=signinInput
	fmt.Println("Name is ",signinInput,"password is ",password)
	if userDao.CheckPassWord(password) ==true  {
		//检验帐号是否激活 todo
		token:=jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256,jwt_lib.MapClaims{
			"id":user.ID,
		})
		//mysecret写死，可以自由设置但要和middler保持一致
		tokenString,err:=token.SignedString([]byte("mysecret"))
		if err!=nil{
			fmt.Println(err.Error())
			SendErrJSON("内部错误.",c)
			return 
		}
	 fmt.Println("user token is ",tokenString)
		c.SetCookie("token",tokenString,100,"/","",true,true)
		c.JSON(http.StatusOK,gin.H{
			"errNo":200,
			"msg":"success",
			"data":gin.H{
				"token":tokenString,
				"user":user,
			},
		})
		return 
	}
	SendErrJSON("帐号或密码错误",c)
}
//signup 用户注册
func Signup(c *gin.Context){
	fmt.Println("signup 用户注册")
	SendErrJSON:=common.SendErrJSON
	type UserReqData struct {
		Name string `json:"name" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}
	var userData UserReqData
	if err:=c.ShouldBindWith(&userData,binding.JSON);err!=nil{
		SendErrJSON("参数无效",c)
		return 
	}		
	var user *model.User
	userDao:=&dao.User{}
	user.Pass=dao.EncryptPassword(userData.Password,userDao.Salt())
	user.Name=userData.Name
	user.Create=time.Now()
	
	if err:=dao.SaveUserMysql(user);err!=nil{
		SendErrJSON("error",c)
		fmt.Println("保存数据出错")
		return 
	}

	if err:=dao.UserToRedis(user);err!=nil{
		fmt.Println("filed redis store")
	}
	c.JSON(http.StatusOK,gin.H{
		"errNo":200,
		"msg":"success",
		"data":user,
	})
}
//用户退出
func Signout(c *gin.Context) {
	fmt.Println("用户退出")
		
}
//修改用户资料




