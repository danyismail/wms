package controllers

import (
	"github.com/astaxie/beego"
	Users "wms/repo/users"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) All() {
	
	repo := Users.UsersRepo{}

	result, err := repo.GetAll()
	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "users.html"
	}

	c.Data["list"] = result
	c.Layout = "template.html"
	c.TplName = "users.html"	
}


// func (c *MainController) Delete(){
	
// 	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
// 	o := orm.NewOrm()
// 	o.Using("default")

// 	if num, err := o.Delete(&models.User{Id:userId}); err == nil {
// 		beego.Info("Record Deleted. ", num)
// 		c.Data["num"] = num
// 		c.Layout = "template.html"
// 		c.TplName = "users.html"
// 	} else {
// 		beego.Error("Record couldn't be deleted. Reason: ", err)
// 		c.Data["err"] = err
// 		c.Layout = "template.html"
// 		c.TplName = "users.html"
// 	}
// }