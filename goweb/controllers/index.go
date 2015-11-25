package controllers

import (
	"github.com/astaxie/beego"
	log "github.com/cihub/seelog"
	"github.com/nivance/go-example/goweb/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	log.Info("Getall")
	this.Data["blogs"] = models.GetAll()
	log.Info(this.Data["blogs"])
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}
