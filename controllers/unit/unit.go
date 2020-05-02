package controllers

import (
	"github.com/astaxie/beego"
	Units "webgudang/repo/units"
	"github.com/astaxie/beego/orm"
)

type UnitController struct {
	beego.Controller
}

func (c *UnitController) All() {	
	o := orm.NewOrm()
	o.Using("default")
	
	repo := Units.UnitRepo{}
	result, err := repo.GetAll()

	if err != nil {

	}

	c.Data["units"] = result
	c.Layout = "template.html"
	c.TplName = "unit.html"
	
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