package middleware

import (
	"github.com/gin-gonic/gin"
	config2 "hechuangfil-admin/config"
	"hechuangfil-admin/handler"
	jwt "hechuangfil-admin/pkg/jwtauth"
	"strings"
	"time"
)

func AuthInit() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour * 12,
		MaxRefresh:      time.Hour * 12 ,
		IdentityKey:     config2.ApplicationConfig.JwtSecret,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

}

func CheckLANIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ips := strings.Split(ctx.ClientIP(),".")
		if ips[0]!="10"&&ips[0]!="172"&&ips[0]!="192"&&ips[0]!="127"{
			ctx.String(403, "quan xian cuo wu")
			ctx.Abort()
			return
		}
		//log.Println("ctx.ClientIP():",ctx.ClientIP())
		ctx.Next()
	}
}