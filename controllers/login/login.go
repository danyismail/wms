package controllers

import (
	_ "wms/models"
	"github.com/astaxie/beego"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "login.html" //buat load halaman
}
