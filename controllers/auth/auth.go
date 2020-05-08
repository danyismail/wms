package controllers

import(
	"wms/models"
	Auth "wms/repo/auth"
	"github.com/astaxie/beego"
)

type AuthController struct{
	beego.Controller
}


func(c *AuthController) Login(){
	repo := Auth.AuthRepo{}
	form := models.LoginForm{}
	form.Email = c.GetString("email")
	form.Password = c.GetString("password")
	result := repo.Login(form)
	if !result  {
		c.Data["errormsg"] = "Login gagal"
		c.Layout = "template.html"
		c.Redirect("/", 302)
	}
	c.Data["errormsg"] = "Login Berhasil"
	c.Layout = "template.html"
	c.Redirect("/dashboard", 302)
}



