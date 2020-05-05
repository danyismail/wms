package controllers

import (
	"github.com/astaxie/beego"
	Dashboard "wms/repo/dashboard"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get(){
	repo := Dashboard.RepoDashboard{}

	totalIn, errTotalIn := repo.GetTotalIncomingGoods()

	if errTotalIn != nil {
		c.Data["errormsg"] = errTotalIn
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	totalOut, errTotalOut := repo.GetTotalOutcomingGoods()

	if errTotalOut != nil {
		c.Data["errormsg"] = errTotalOut
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	totalUser, errTotalUser := repo.GetTotalOutcomingGoods()

	if errTotalUser != nil {
		c.Data["errormsg"] = errTotalUser
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	c.Data["jumlahMasuk"] = totalIn
	c.Data["jumlahKeluar"] = totalOut
	c.Data["jumlahUser"] = totalUser
	c.Layout = "template.html"
	c.TplName = "dashboard.html" //buat load halaman
}


