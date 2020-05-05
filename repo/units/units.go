package repo

import "wms/database"
import "wms/models"
import "github.com/astaxie/beego"
// import "github.com/astaxie/beego/orm"
// import "fmt"

type UnitRepo struct{
	beego.Controller
}

func (repo *UnitRepo) GetAll() (result []models.Units ,err error){
	
    var units []models.Units
	
	if result := database.DB.Find(&units);  result.Error != nil {
		return units, result.Error
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



