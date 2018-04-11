package jwt

import(
	"github.com/gin-gonic/gin"
	dao "github.com/zkc123/dao"
	// "errors"
)
func RefreshTokenCookie(c *gin.Context)  {
	tokenString,err:=c.Cookie("token")
	if tokenString!=""&& err ==nil{
		c.SetCookie("token",tokenString,100,"/","",true,true)
		if user,err:=getUser(c);err==nil{
				dao.UserToRedis(&user)
			}		
		}
		c.Next()
}