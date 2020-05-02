package controllers

import (
	"github.com/astaxie/beego"
	Dashboard "webgudang/repo/dashboard"
	"fmt"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get(){
	repo := Dashboard.RepoDashboard{}

	result, err := repo.GetAll()

	if err != nil {
		fmt.Println("ada error")
	}

	totalIn, errTotalIn := repo.GetTotalIncomingGoods()

	if errTotalIn != nil {
		panic(errTotalIn)
	}

	totalOut, errTotalOut := repo.GetTotalOutcomingGoods()

	if errTotalOut != nil {
		panic(errTotalOut)
	}

	totalUser, errTotalUser := repo.GetTotalOutcomingGoods()

	if errTotalUser != nil {
		panic(errTotalOut)
	}

	c.Data["jumlahMasuk"] = totalIn
	c.Data["jumlahKeluar"] = totalOut
	c.Data["jumlahUser"] = totalUser
	c.Data["list"] = result
	c.Layout = "template.html"
	c.TplName = "dashboard.html" //buat load halaman
}


