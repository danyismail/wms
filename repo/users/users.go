package repo

import "webgudang/models"
import "github.com/astaxie/beego"
import "github.com/astaxie/beego/orm"
import "fmt"

type UsersRepo struct{
	beego.Controller
}

func (repo *UsersRepo) GetAll() (result []models.Users ,err error){
	o := orm.NewOrm()
    o.Using("default")

    var users []models.Users
    num, err := o.QueryTable("users").All(&users) //select * from incoming

    if err != orm.ErrNoRows && num > 0 {
		fmt.Println("ada error", err)
	}
		
		//Handle error
	if err != nil {
		panic (err)
	}
	
	return users, nil
}	


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



