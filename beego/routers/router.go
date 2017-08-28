package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/theplant/auth_examples/beego/conf"
	"github.com/theplant/auth_examples/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Any("/auth/*", func(ctx *context.Context) {
		conf.Auth.NewServeMux().ServeHTTP(ctx.ResponseWriter, ctx.Request)
	})
}
