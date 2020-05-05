package repo

import "wms/database"
import "wms/models"
import "github.com/astaxie/beego"

type UsersRepo struct{
	beego.Controller
}

func (repo *UsersRepo) GetAll() (result []models.Users ,err error){

    var users []models.Users
	if result := database.DB.Find(&users); result.Error != nil {
		return users, result.Error
	}

	return users, nil
}	

/*
func (repo *UsersRepo) GetById(id int) (result models.Incoming ,err error){
	o := orm.NewOrm()
	o.Using("default")
	
	DetailBarang := models.Incoming{}

	err = o.QueryTable("incoming").Filter("id", id).One(&DetailBarang)
	if err == orm.ErrMultiRows {
		// Have multiple records
		fmt.Printf("Returned Multi Rows Not One")
		}
		if err == orm.ErrNoRows {
		fmt.Printf("Not row found")
		}
	return DetailBarang, nil
}	
*/



