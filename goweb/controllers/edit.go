package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nivance/go-example/basic/logs"
	"github.com/nivance/go-example/goweb/models"
	"strconv"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplNames = "edit.tpl"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Entity
	logs.Logger.Info("inputs:", inputs)
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}
