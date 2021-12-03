package routers

import (
	"fzuhelper_launch_screen/middleware/jwt"
	"fzuhelper_launch_screen/routers/api"
	"github.com/gin-gonic/gin"
)

// InitRouter 载入路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/example/hello", api.Hello)
	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{

	}
	return r
}
