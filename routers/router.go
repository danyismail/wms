package routers

import (
	Login 		"webgudang/controllers/login"
	Api 		"webgudang/controllers/api"
	Dashboard 	"webgudang/controllers/dashboard"
	Users 		"webgudang/controllers/users"
	Unit 		"webgudang/controllers/unit"
	Incoming  	"webgudang/controllers/incoming"
	Outcoming 	"webgudang/controllers/outcoming"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &Login.MainController{})
	beego.Router("/goods", &Api.MainController{}, "*:View")
	beego.Router("/dashboard", &Dashboard.MainController{})
	beego.Router("/users/delete/:id([0-9]+)", &Users.UserController{}, "*:Delete")
	beego.Router("/users", &Users.UserController{}, "*:View")
	
	beego.Router("/unit", &Unit.UnitController{}, "*:View")
	
	beego.Router("/incoming", &Incoming.IncomingController{})
	
	beego.Router("/outcoming", &Outcoming.OutcomingController{}, "*:View")
}

