package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (h *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", h.Hello)
}

//解析/hello
func (h *HelloController) Hello(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "你好啊",
		"name":    "是啊",
	})
}
