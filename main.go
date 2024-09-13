package main

import (
	"gin-blog/core"
	"gin-blog/flag"
	"gin-blog/global"
	"gin-blog/routers"
)

func main() {
	// 初始化配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.InitGorm()
	// 命令行参数绑定，数据迁移
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
	}
	// 初始化路由
	r := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("启动成功,启动地址为:%s", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatalf("启动失败，错误信息：%s", err.Error())
	}
}
