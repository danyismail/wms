package database

import(
	_ "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var DB *gorm.DB

func init(){
	fmt.Println("koneksi database terpanggil")
	db, err := gorm.Open("mysql", "root:root@(localhost:8889)/web_gudang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gagal koneksi ke database")
	} else {
		DB = db
		DB.SingularTable(true) //Ini config dari GORM
		fmt.Println("koneksi ke database sukses")
	}
}

