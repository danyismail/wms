package repo

import "webgudang/models"
import "github.com/astaxie/beego"
import "github.com/astaxie/beego/orm"
import "fmt"

type GoodsRepo struct{
	beego.Controller
}

func (repo *GoodsRepo) GetAll() (result []models.Incoming ,err error){
	o := orm.NewOrm()
    o.Using("default")

    var barangMasuk []models.Incoming
    num, err := o.QueryTable("incoming").All(&barangMasuk) //select * from incoming

    if err != orm.ErrNoRows && num > 0 {
		return barangMasuk, nil
	}
	
	return barangMasuk, err
}	


func (repo *GoodsRepo) GetById(id int) (result models.Incoming ,err error){
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


func (repo *GoodsRepo) Add(form models.Incoming) (bool, error){
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(&form)
	if err != nil {
		fmt.Println("ada error di repo", err)
		return false, err
	}
	fmt.Println(id)
	return true, nil
}	










