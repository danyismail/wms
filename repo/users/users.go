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


