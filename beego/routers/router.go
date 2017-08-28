package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth_themes/clean"
	"github.com/theplant/auth_examples/beego/controllers"
)

var (
	db, _     = gorm.Open("sqlite3", "test.db")
	Auth      = clean.New(&auth.Config{DB: db})
	Authority = authority.New(&authority.Config{
		Auth: Auth,
		RedirectPathAfterAccessDenied: "/auth/login",
	})
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Any("/auth/*", func(ctx *context.Context) {
		Auth.NewServeMux().ServeHTTP(ctx.ResponseWriter, ctx.Request)
	})
}
