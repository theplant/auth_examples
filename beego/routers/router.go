package routers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/middlewares"
	"github.com/theplant/auth_examples/beego/conf"
	"github.com/theplant/auth_examples/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Any("/auth/*", func(ctx *context.Context) {
		conf.Auth.NewServeMux().ServeHTTP(ctx.ResponseWriter, ctx.Request)
	})

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		middlewares.Apply(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx.Request = req
		})).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	})
}
