package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nivance/go-example/basic/logs"
	"github.com/nivance/go-example/goweb/models"
	"time"
)

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.tpl"
	this.TplNames = "new.tpl"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var blog models.Entity
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	logs.Logger.Info("blog:", blog)
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
