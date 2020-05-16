package controllers

import (
	"wms/models"
	UserRepo "wms/repo/users"
	"github.com/astaxie/beego"
	"fmt"
)


type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html" //buat load halaman
}

func (c *LoginController) Logout() {
	c.DelSession("controllerSession")
	c.Redirect("/", 302)
}

func (c *LoginController) Login() {
	repo := UserRepo.UsersRepo{}
	userModel :=  models.LoginForm{}
	//mau nampung data dari form
	userModel.Email 	= c.GetString("email")
	userModel.Password 		= c.GetString("password")
	
	result := repo.Login(userModel)
	if result == "gagal login" {
		fmt.Println(result)
		c.Redirect("/",302)
	} else {
		c.SetSession("controllerSession", int(1))
        c.Data["sessionID"] = "191919"
		c.Data["messageSuccess"] = "Berhasil login"
		c.Redirect("/dashboard",302)
	}
}
