package database

import(
	"github.com/astaxie/beego/orm"
	"webgudang/models"
	_ "github.com/go-sql-driver/mysql"  
)

// func init(){
// 	orm.RegisterDriver("mysql", orm.DRMySQL)
//     // parts := []string{"root", ":", "","@tcp(127.0.0.1:", "3306)/", "web_gudang"}
// 	// orm.RegisterDataBase("default", "mysql", "root/orm_test?charset=utf8")
// 	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:8889)/web_gudang")
// 	orm.RegisterModel(new(models.User))
// 	orm.RegisterModel(new(models.Unit))
// 	orm.RegisterModel(new(models.Incoming))
// 	orm.RegisterModel(new(models.Outcoming))
// }

func Conn(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
    // parts := []string{"root", ":", "","@tcp(127.0.0.1:", "3306)/", "web_gudang"}
	// orm.RegisterDataBase("default", "mysql", "root/orm_test?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:8889)/web_gudang")
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Unit))
	orm.RegisterModel(new(models.Incoming))
	orm.RegisterModel(new(models.Outcoming))
}