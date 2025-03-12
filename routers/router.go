package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/ccppoo/f1-api/docs"
	"github.com/ccppoo/f1-api/pkg/jwt"
	"github.com/ccppoo/f1-api/routers/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	JWTMiddleware := jwt.GetJWTMiddleware()
	HandlerMiddleware := jwt.HandlerMiddleware(JWTMiddleware)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/login", JWTMiddleware.LoginHandler)
	authRoute := r.Group("/auth")
	authRoute.GET("/refresh_token", JWTMiddleware.RefreshHandler)
	authRoute.POST("/logout", JWTMiddleware.LogoutHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	authedRoutes := r.Group("/authed")
	authedRoutes.Use(HandlerMiddleware)
	{
		authedRoutes.GET("", api.HelloWorld)
	}

	return r
}
