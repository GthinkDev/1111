package api

import (
	"gin-blog/config"
	"gin-blog/core"
	"gin-blog/global"
	"gin-blog/models/res"
	"github.com/gin-gonic/gin"
)

func (s SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
}

func (s SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	cr := config.Email{}
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParams, c)
		return
	}
	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
