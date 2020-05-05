package controllers

import (
	"github.com/astaxie/beego"
	Units "wms/repo/units"
)

type UnitController struct {
	beego.Controller
}

func (c *UnitController) All() {	
	repo := Units.UnitRepo{}
	result, err := repo.GetAll()

	if err != nil {
		c.Data["units"] = err
		c.Layout = "template.html"
		c.TplName = "unit.html"
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