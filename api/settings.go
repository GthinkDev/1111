package api

import (
	"gin-blog/config"
	"gin-blog/core"
	"gin-blog/global"
	"gin-blog/models/res"

	"github.com/gin-gonic/gin"
)

type SettingsApi struct {
}

func (s SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}

func (s SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	cr := config.SiteInfo{}
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParams, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
