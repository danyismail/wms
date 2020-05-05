package database

import(
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
//   "wms/models"
  "fmt"
)

var DB *gorm.DB

func init(){
	fmt.Println("koneksi database terpanggil")
	/*
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:8889)/web_gudang")
	orm.RegisterModel(new(models.User))  //tabel user
	orm.RegisterModel(new(models.Unit))
	orm.RegisterModel(new(models.Incoming))
	orm.RegisterModel(new(models.Outcoming))
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Units))
	*/
	db, err := gorm.Open("mysql", "root:root@(localhost:8889)/web_gudang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gagal konek database")
	}   else {
		DB = db
		DB.SingularTable(true)
		fmt.Println("konek database sukses")
	}
}

