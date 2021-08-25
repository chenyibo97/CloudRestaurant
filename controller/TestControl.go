package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studygo2/CloudRestaurant/tool"
)

type TestController struct {
}

func (h *TestController) Router(engine *gin.Engine) {
	engine.POST("/api/redistest", h.RedisTest)
	engine.GET("/api/redistest", h.RedisTest)
	engine.POST("/api/upload/redis", h.RedisTest)
	engine.POST("/api/upload/redis2", h.RedisTest2)
	engine.GET("/api/upload/redis2", h.RedisTest2)

	engine.GET("/test", h.nexttest1, h.nexttest2)
}

func (h *TestController) RedisTest(ctx *gin.Context) {
	tool.SetSession(ctx, 1, 2)
	session := tool.GetSession(ctx, 1)
	fmt.Println(session)
	ctx.JSON(200, session)
}
func (h *TestController) RedisTest2(ctx *gin.Context) {
	//tool.SetSession(ctx,1,2)
	session := tool.GetSession(ctx, 1)
	fmt.Println(session)
	ctx.JSON(200, session)
}
func (h *TestController) nexttest1(ctx *gin.Context) {
	fmt.Println("1")
}
func (h *TestController) nexttest2(ctx *gin.Context) {
	fmt.Println("2")
}
