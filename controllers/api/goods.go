package controllers

import (
	"github.com/astaxie/beego"
	"webgudang/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetIncomingGoods() {
	o := orm.NewOrm()
    o.Using("default")

    var barangMasuk []*models.Incoming
    num, err := o.QueryTable("incoming").All(&barangMasuk) //select * from incoming

    if err != orm.ErrNoRows && num > 0 {
		_, errJ := json.Marshal(barangMasuk) //Konversi dari struct ke string
		//Handle error
		if errJ != nil {
			panic (errJ)
		}
	
	c.Data["json"] = barangMasuk
	c.ServeJSON()
	}
}

func (c *MainController) GetOutcomingGoods() {
	o := orm.NewOrm()
    o.Using("default")

    var barangKeluar []*models.Outcoming
    num, err := o.QueryTable("outcoming").All(&barangKeluar) //select * from incoming

    if err != orm.ErrNoRows && num > 0 {
		_, errJ := json.Marshal(barangKeluar) //Konversi dari struct ke string
		//Handle error
		if errJ != nil {
			panic (errJ)
		}
	
	c.Data["json"] = barangKeluar
	c.ServeJSON()
	}
}

func (c *MainController) InDetail() {
	o := orm.NewOrm()
	o.Using("default")
	
	var detailBarangMasuk models.Incoming
	id := c.Ctx.Input.Param(":id")
	i ,_  := strconv.Atoi(id)
	err := o.QueryTable("incoming").Filter("id", i).One(&detailBarangMasuk)
	if err == orm.ErrMultiRows {
    // Have multiple records
    fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
    fmt.Printf("Not row found")
	}
	
	c.Data["json"] = detailBarangMasuk
	c.ServeJSON()	
	
}

func (c *MainController) OutDetail() {
	o := orm.NewOrm()
	o.Using("default")
	
	var detailBarangKeluar models.Outcoming
	id := c.Ctx.Input.Param(":id")
	i ,_  := strconv.Atoi(id)
	err := o.QueryTable("outcoming").Filter("id", i).One(&detailBarangKeluar)
	if err == orm.ErrMultiRows {
    // Have multiple records
    fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
    fmt.Printf("Not row found")
	}
	c.Data["data"] = detailBarangKeluar
	c.TplName = "detail.html"
}

func (c *MainController) Delete() {
	o := orm.NewOrm()
	o.Using("default")
	
	var detailBarangKeluar models.Outcoming
	id := c.Ctx.Input.Param(":id")
	i ,_  := strconv.Atoi(id)
	err := o.QueryTable("outcoming").Filter("id", i).One(&detailBarangKeluar)
	if err == orm.ErrMultiRows {
    // Have multiple records
    fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
    fmt.Printf("Not row found")
	}
	c.Data["data"] = detailBarangKeluar
	c.TplName = "detail.html"
}