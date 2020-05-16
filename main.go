package main

import (
	_ "wms/routers"
	"github.com/astaxie/beego" 
	// "github.com/astaxie/beego/orm"
)



func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.Session.SessionOn = true
		//Put XSRF 
		// beego.BConfig.WebConfig.EnableXSRF = true
		// beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
		// beego.BConfig.WebConfig.XSRFExpire = 3600
	}


	
	beego.Run()
}

