package routers

import (
	Login 		"webgudang/controllers/login"
	Api 		"webgudang/controllers/api"
	Incoming 		"webgudang/controllers/incoming"
	Outcoming 		"webgudang/controllers/outcoming"
	Dashboard 	"webgudang/controllers/dashboard"
	Users 		"webgudang/controllers/users"
	Unit 		"webgudang/controllers/unit"
	"github.com/astaxie/beego"
)

func init() {

	//GOODS ROUTER
	beego.Router("/", &Login.MainController{})
	beego.Router("/dashboard", &Dashboard.MainController{})
	beego.Router("/incoming-goods",&Incoming.MainController{}, "get:All")
	beego.Router("/incoming-goods/add",&Incoming.MainController{}, "post:Add")
	beego.Router("/incoming-goods/detail/:id([0-9]+)", &Incoming.MainController{}, "get:Detail")
	
	beego.Router("/outcoming-goods", &Outcoming.MainController{})
	beego.Router("/outcoming-goods/detail", &Outcoming.MainController{})

	//UNITS ROUTER
	beego.Router("/units", &Unit.UnitController{}, "get:All")

	//USERS ROUTER
	beego.Router("/users", &Users.UserController{}, "get:All")

	//API Controllers
	beego.Router("/api", &Dashboard.MainController{})
	beego.Router("/api/incoming", &Api.MainController{}, "*:GetIncomingGoods")
	beego.Router("/api/incoming", &Api.MainController{}, "*:GetIncomingGoods")
	
	// beego.Router("/backend/incoming/detail/:id([0-9]+)", &Api.MainController{}, "get:InDetail")
	
	beego.Router("/api/outcoming", &Api.MainController{}, "*:GetOutcomingGoods")
	beego.Router("/api/outcoming/detail/:id([0-9]+)", &Api.MainController{}, "*:OutDetail")
	beego.Router("/api/outcoming/update/:id([0-9]+)", &Api.MainController{}, "*:OutDetail")
	
	beego.Router("/users/delete/:id([0-9]+)", &Users.UserController{}, "*:Delete")
	beego.Router("/users", &Users.UserController{}, "get:All")
	

}

