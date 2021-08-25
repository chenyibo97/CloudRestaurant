package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(engine *gin.Engine) {
	cofig := GetCofig().Redis
	store, err := redis.NewStore(10, "tcp", cofig.Addr+":"+cofig.Port, "", []byte("secret"))
	if err != nil {
		fmt.Println("创建redis store失败", err)
	}

	engine.Use(sessions.Sessions("mysession", store))

}
func SetSession(ctx *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(ctx)

	if session == nil {
		return nil
	}
	session.Set(key, value)
	fmt.Println("session:", session)
	fmt.Println("redis设置key:", key, "value:", value)
	return session.Save()
}

func GetSession(context *gin.Context, key interface{}) (value interface{}) {

	session := sessions.Default(context)
	fmt.Println(session)
	fmt.Println("redis获取key:", key, "value:", session.Get(key))
	return session.Get(key)
}
