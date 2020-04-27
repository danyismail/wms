package controllers

import (
	_ "webgudang/models"
	"github.com/astaxie/beego"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//Layout <- Template
	// c.Layout = "template.html"
	c.TplName = "login.html" //buat load halaman
}
