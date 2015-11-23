package routers

import (
	"github.com/nivance/rebotweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
