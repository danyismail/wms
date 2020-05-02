package main

import (
	_ "webgudang/routers"
	"github.com/astaxie/beego" 
	"github.com/astaxie/beego/orm"
	"webgudang/database"
)



func main() {

	database.Conn()

	if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        // beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	}

	orm.Debug = true
	beego.Run()
}

