package repo

import "wms/models"
import "wms/database"
import "github.com/astaxie/beego"
// import "github.com/astaxie/beego/orm"
// import "fmt"

type GoodsRepo struct{
	beego.Controller
}

//Buat kondisi where jumlah tidak != 0
func (repo *GoodsRepo) GetAll() (result []models.Incoming ,err error){
	// o := orm.NewOrm()
    // o.Using("default")

    var barangMasuk []models.Incoming
	// num, err := o.QueryTable("incoming").All(&barangMasuk) //select * from incoming
	if result := database.DB.Find(&barangMasuk); result.Error != nil {
		return barangMasuk, result.Error
	}

	return barangMasuk, err
}	


func (repo *GoodsRepo) GetById(id int) (result models.Incoming ,err error){
	
	DetailBarang := models.Incoming{}

	if result := database.DB.Where("id = ?", id).Find(&DetailBarang); result.Error != nil {
		return DetailBarang , result.Error
	}

	return DetailBarang , nil
}	



func (repo *GoodsRepo) Delete(id int) (string ,error){
	DeleteBarang := models.Incoming{}
	DeleteBarang.Id = id
	if result := database.DB.Delete(&DeleteBarang); result.Error != nil {
		return "Hapus barang gagal", result.Error
	}
	
	return "Delete barang berhasil", nil
}	



func (repo *GoodsRepo) Add(form models.Incoming) ( bool, error){
	
	database.DB.Create(&form)
	return true, nil
}	

func (repo *GoodsRepo) OutProccess(form models.Outcoming) ( string, error){
	if result := database.DB.Create(&form); result.Error != nil {
		return "proses barang gagal", result.Error
	}
	return "proses barang berhasil", nil
}	

func (repo *GoodsRepo) Update(form models.Incoming) (string, error){
	UpdateBarang := models.Incoming{}
	if result := database.DB.Model(&UpdateBarang).Updates(form); result.Error != nil {
		return "Update gagal", result.Error
	}
	return "Update data berhasil", nil
}











