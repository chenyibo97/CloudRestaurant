package tool

import "github.com/gin-gonic/gin"

const (
	SUCESS int = 0
	Fail   int = 1
)

func Sucess(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code": SUCESS,
		"msg":  "成功",
		"data": v,
	})
}
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code": Fail,
		"msg":  "失败",
		"data": v,
	})
}
