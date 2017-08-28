package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qor/auth/claims"
	"github.com/theplant/auth_examples/beego/conf"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	Claims, _ := conf.Auth.Get(c.Ctx.Request)
	if Claims == nil {
		Claims = &claims.Claims{}
	}

	c.Data["LoggedAt"] = Claims.LastLoginAt
	c.Data["LongestDistraction"] = Claims.LongestDistractionSinceLastLogin
	c.Data["CurrentPath"] = c.Ctx.Request.URL.Path
	c.TplName = "index.tpl"
}
