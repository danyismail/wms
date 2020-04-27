package controllers

import (
	"github.com/astaxie/beego"
	"webgudang/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type OutcomingController struct {
	beego.Controller
}

func (c *OutcomingController) View() {
	o := orm.NewOrm()
    o.Using("default")

    var outcomingGoods []*models.Outcoming
    num, err := o.QueryTable("outcoming").All(&outcomingGoods)

    if err != orm.ErrNoRows && num > 0 {
		out, errJ := json.Marshal(outcomingGoods) //Konversi dari struct ke string
		//Handle error
		if errJ != nil {
			panic (errJ)
		}
		c.Data["records"] = string(out)
		c.Layout = "template.html"
		c.TplName = "outcoming.html"
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