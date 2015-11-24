package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nivance/go-example/goweb/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}
