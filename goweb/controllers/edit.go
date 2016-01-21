package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nivance/go-example/basic/logs"
	"github.com/nivance/go-example/goweb/models"
	"strconv"
	"time"
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
	dataStr := inputs.Get("created")
	date, err := time.Parse("2006-01-02 15:04:05", dataStr)
	if err != nil {
		logs.Logger.Warn("time parse error:", dataStr)
		date = time.Now()
	}
	blog.Created = date
	models.UpdateBlog(blog)
	this.Ctx.Redirect(302, "/")
}
