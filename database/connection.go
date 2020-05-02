package database

import(
	"github.com/astaxie/beego/orm"
	"webgudang/models"
	_ "github.com/go-sql-driver/mysql"
)

func Conn(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:8889)/web_gudang")
	orm.RegisterModel(new(models.User))  //tabel user
	orm.RegisterModel(new(models.Unit))
	orm.RegisterModel(new(models.Incoming))
	orm.RegisterModel(new(models.Outcoming))
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Units))
}