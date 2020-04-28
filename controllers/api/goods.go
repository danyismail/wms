package controllers

import (
	"github.com/astaxie/beego"
	"webgudang/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) View() {
	o := orm.NewOrm()
    o.Using("default")

    var barangMasuk []*models.Incoming
    num, err := o.QueryTable("incoming").All(&barangMasuk)

    if err != orm.ErrNoRows && num > 0 {
		_, errJ := json.Marshal(barangMasuk) //Konversi dari struct ke string
		//Handle error
		if errJ != nil {
			panic (errJ)
		}
	c.Data["json"] = barangMasuk
	// c.Layout = "template.html"
	// c.TplName = "incoming.html"
	c.ServeJSON()
	}
}