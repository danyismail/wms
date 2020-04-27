package api

import (
	"github.com/astaxie/beego"
)


type IncomingController struct {
	beego.Controller
}

func (c *IncomingController) Get() {
	//Layout <- Template
	c.Layout = "template.html"
	c.TplName = "incoming.html" //buat load halaman
}
