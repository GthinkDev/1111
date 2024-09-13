package routers

import (
	"gin-blog/api"
)

func (r *RouterGroup) SettingsRouter() {
	settings := api.ApiGroup.Settings
	r.GET("settings/:name", settings.SettingsInfoView)
	r.PUT("settings/:name", settings.SettingsInfoUpdateView)
}
