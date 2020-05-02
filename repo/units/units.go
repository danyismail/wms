package repo

import "webgudang/models"
import "github.com/astaxie/beego"
import "github.com/astaxie/beego/orm"
import "fmt"

type UnitRepo struct{
	beego.Controller
}

func (repo *UnitRepo) GetAll() (result []models.Units ,err error){
	o := orm.NewOrm()
    o.Using("default")

    var units []models.Units
    num, err := o.QueryTable("units").All(&units) //select * from incoming

    if err != orm.ErrNoRows && num > 0 {
		fmt.Println(err)
	}
	return units, nil
}


// func (repo *UsersRepo) GetById() (result []models.Incoming ,err error){
// 	o := orm.NewOrm()
// 	o.Using("default")
	
// 	DetailBarang := models.Incoming{}

// 	err = o.QueryTable("incoming").Filter("id", id).One(&DetailBarang)
// 	if err == orm.ErrMultiRows {
// 		// Have multiple records
// 		fmt.Printf("Returned Multi Rows Not One")
// 		}
// 		if err == orm.ErrNoRows {
// 		fmt.Printf("Not row found")
// 		}
// 	return DetailBarang, nil
// }	



