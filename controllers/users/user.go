package controllers

import (
	"github.com/astaxie/beego"
	// "strconv"
	"webgudang/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) View() {
	o := orm.NewOrm()
    o.Using("default")

    var arrUser []*models.User
    num, err := o.QueryTable("user").All(&arrUser)

    if err != orm.ErrNoRows && num > 0 {
		out, errJ := json.Marshal(arrUser) //Konversi dari struct ke string
		//Handle error
		if errJ != nil {
			panic (errJ)
		}
	c.Data["users"] = string(out)
	c.Layout = "template.html"
	c.TplName = "users.html"
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
// 		c.TplName = "users.html"
// 	} else {
// 		beego.Error("Record couldn't be deleted. Reason: ", err)
// 		c.Data["err"] = err
// 		c.Layout = "template.html"
// 		c.TplName = "users.html"
// 	}
// }