package main

import (
	_ "wms/routers"
	"github.com/astaxie/beego" 
	// "github.com/astaxie/beego/orm"
)



func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	beego.Run()
}

