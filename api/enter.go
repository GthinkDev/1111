package api

type ApiGroups struct {
	Settings SettingsApi
}

var ApiGroup = new(ApiGroups)

// InitRouter 初始化路由
func InitRouter() {
}
