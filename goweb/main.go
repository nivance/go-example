package main

import (
	"github.com/astaxie/beego"
	"github.com/nivance/go-example/basic/logs"
	"github.com/nivance/go-example/goweb/controllers"
)

func main() {

	logs.Logger.Debug("-------router is starting-------")
	//显示博客首页
	beego.Router("/", &controllers.IndexController{})
	//查看博客详细信息
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})
	//新建博客博文
	beego.Router("/new", &controllers.NewController{})
	//	//删除博文
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//编辑博文
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})
	logs.Logger.Info("------goweb is started")
	beego.Run()
}
