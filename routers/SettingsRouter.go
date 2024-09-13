package routers

import (
	"gin-blog/api"
)

func (r *RouterGroup) SettingsRouter() {
	settings := api.ApiGroup.Settings
	r.GET("/settings", settings.SettingsInfoView)
	r.PUT("/settings", settings.SettingsInfoUpdateView)
	r.GET("/email", settings.SettingsEmailInfoView)
	r.PUT("/email", settings.SettingsEmailInfoUpdateView)
}
