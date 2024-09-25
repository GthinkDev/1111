package api

type ApiGroups struct {
	Settings SettingsApi
	Images   ImagesApi
}

var ApiGroup = new(ApiGroups)

// InitRouter 初始化路由
func InitRouter() {
}
