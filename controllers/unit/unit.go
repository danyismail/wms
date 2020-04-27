package controllers

import (
	"github.com/astaxie/beego"
	// "strconv"
	"webgudang/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type UnitController struct {
	beego.Controller
}

func (c *UnitController) View() {	
	o := orm.NewOrm()
    o.Using("default")

    var arrUnit []*models.Unit
    num, err := o.QueryTable("unit").All(&arrUnit)

    if err != orm.ErrNoRows && num > 0 {
		out, errJ := json.Marshal(arrUnit) //Konversi dari struct ke string
		//Handle error
		if errJ != nil {
			panic (errJ)
		}
	c.Data["unit"] = string(out)
	c.Layout = "template.html"
	c.TplName = "unit.html"
	}
	
}


// func (c *MainController) Delete(){
	
// 	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
// 	o := orm.NewOrm()
// 	o.Using("default")

// 	if num, err := o.Delete(&models.User{Id:userId}); err == nil {
// 		beego.Info("Record Deleted. ", num)
// 		c.Data["num"] = num
// 		c.Layout = "template.html"
// 		c.TplName = "dashboard.html"
// 	} else {
// 		beego.Error("Record couldn't be deleted. Reason: ", err)
// 		c.Data["err"] = err
// 		c.Layout = "template.html"
// 		c.TplName = "users.html"
// 	}
// }