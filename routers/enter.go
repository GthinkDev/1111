package routers

import (
	"gin-blog/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) // 设置gin的模式 去掉启动时 gin 默认的一坨信息
	r := gin.Default()

	// 路由组
	api := r.Group("/api")

	// 版本
	v1 := api.Group("/v1")

	// 注册路由-系统配置 Api
	routerGroup := RouterGroup{v1}
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()

	return r
}
