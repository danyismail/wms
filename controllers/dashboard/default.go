package controllers

import (
	"github.com/astaxie/beego"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//Layout <- Template
	c.Layout = "template.html"
	c.TplName = "dashboard.html" //buat load halaman
}
