package routers

import (
	"github.com/astaxie/beego"
	"spider.ohmgood.com/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
