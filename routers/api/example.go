package api

import (
	"fzuhelper_launch_screen/pkg/app"
	"fzuhelper_launch_screen/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Hello 路由样例
func Hello(c *gin.Context){
	g := app.Gin{C: c}
	params := struct {
		Name  string `form:"name" json:"name" xml:"name" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&params); err != nil {
		g.Response(http.StatusOK, e.InvalidParams, "参数错误")
		return
	}
	id := c.MustGet("id").(string)
	g.Response(http.StatusOK, e.SUCCESS, "hello " + id + params.Name)
	return
}
