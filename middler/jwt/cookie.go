package jwt

import(
	"github.com/gin-gonic/gin"
	dao "github.com/zkc123/dao"
	"fmt"
	"net/http"
	// "errors"
)
func RefreshTokenCookie(c *gin.Context)  {

	// tokenString,err:=c.Cookie("token")
	var tokenString string
	var err error
	if cookie, err := c.Request.Cookie("token"); err == nil {
		tokenString = cookie.Value
	}
	fmt.Println("刷新token",tokenString)
	if tokenString!=""&& err ==nil{
		cookie:=&http.Cookie{
			Name:"token",
			Value:tokenString,
			MaxAge:1<<31,
			Path:"/",
			Secure:false,
			HttpOnly:false,
		}
		// c.SetCookie()
		http.SetCookie(c.Writer,cookie)
		// c.SetCookie(....)未生效
		fmt.Println("重新设置token",cookie.Name)
		if user,err:=getUser(c);err==nil{
				dao.UserToRedis(&user)
			}		
		}
		c.Next()
}