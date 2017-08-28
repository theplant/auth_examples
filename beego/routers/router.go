package routers

import (
	"github.com/theplant/auth_examples/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
