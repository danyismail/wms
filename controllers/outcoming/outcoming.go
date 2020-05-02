package api

import (
	"github.com/astaxie/beego"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//Layout <- Template
	c.Layout = "template.html"
	c.TplName = "outcoming.html" //buat load halaman
}
