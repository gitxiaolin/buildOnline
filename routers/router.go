package routers

import (
	"github.com/gitxiaolin/buildonline/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
