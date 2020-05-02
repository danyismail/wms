package repo

import "webgudang/models"
import "github.com/astaxie/beego"
import "github.com/astaxie/beego/orm"
import "fmt"

type RepoDashboard struct{
	beego.Controller
}

func (repo *RepoDashboard) GetAll() (result []models.Incoming ,err error){
	o := orm.NewOrm()
    o.Using("default")

    var barangMasuk []models.Incoming
    num, err := o.QueryTable("incoming").All(&barangMasuk) //select * from incoming

    if err != orm.ErrNoRows && num > 0 {

		fmt.Println("ada error", err)

	}
		
		//Handle error
		if err != nil {
			panic (err)
		}
	
	return barangMasuk, nil
}	

func (repo *RepoDashboard) GetTotalIncomingGoods() (result int, err error){
	o := orm.NewOrm()
    o.Using("default")

    var totalBarangMasuk int
    total, err := o.QueryTable("incoming").Count() 

    if err != nil {
		return totalBarangMasuk,err
	}
	totalBarangMasuk = int(total)
	return totalBarangMasuk, nil 
}	

func (repo *RepoDashboard) GetTotalOutcomingGoods() (result int, err error){
	o := orm.NewOrm()
    o.Using("default")

    var totalBarangKeluar int
    total, err := o.QueryTable("outcoming").Count() 

    if err != nil {
		return totalBarangKeluar,err
	}
	totalBarangKeluar = int(total)
	return totalBarangKeluar, nil 
}	

func (repo *RepoDashboard) GetTotalUser() (result int, err error){
	o := orm.NewOrm()
    o.Using("default")

    var totalUser int
    count, err := o.QueryTable("users").Count() 

    if err != nil {
		return totalUser,err
	}
	totalUser = int(count)
	return totalUser, nil 
}	



