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

type Uri struct {
	Name string `uri:"name"` // 路由参数
}

func (s SettingsApi) SettingsInfoView(c *gin.Context) {
	cr := Uri{}
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParams, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("参数错误", c)

	}
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
