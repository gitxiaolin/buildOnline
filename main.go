package main

import (
	"github.com/gitxiaolin/buildonline/controllers"
	_ "github.com/gitxiaolin/buildonline/routers"
	"github.com/astaxie/beego"
	"fmt"
	"os"
	"os/signal"
)

const (
	APP_VER = "localhost"
)

func main() {
	err := os.MkdirAll("/usr/local/src/interim", 0777)
	if err != nil {
		fmt.Println("make interim document faild")
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("received ctrl+c(%v)\n", sig)
			//Linux
			//err := os.RemoveAll("/usr/local/src/interim")
			//Windows
			err ：= os.RemoveAll("D:\\interim")
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		}
	}()
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.Router("/build", &controllers.AppController{})
	fmt.Println("\n\nplease ues `ctrl + c` to stop serving")
	beego.Run()
}
