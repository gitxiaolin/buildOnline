package main

import (
	"buildonline/controllers"
	_ "buildonline/routers"
	"github.com/astaxie/beego"
)

const (
	APP_VER = "localhost"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.Router("/build", &controllers.AppController{})
	beego.Run()
}
