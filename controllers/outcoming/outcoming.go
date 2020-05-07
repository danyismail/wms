package api

import (
	"github.com/astaxie/beego"
	"wms/repo/goods"
	"wms/models"
	"strconv"
	"time"
	"fmt"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//Layout <- Template
	c.Layout = "template.html"
	c.TplName = "outcoming.html" //buat load halaman
}

func (c *MainController) OutProccess() {
	t := time.Now()
	goodsRepo := repo.GoodsRepo{}
	outModel :=  models.Outcoming{}
	//mau nampung data dari form
	getID, _ := strconv.Atoi(c.GetString("id"))
	outModel.Id 	= getID
	outModel.TransaksiId 	= c.GetString("trxId")
	outModel.TanggalMasuk 	= c.GetString("tanggal")
	outModel.TanggalKeluar 	= t.Format("2006/01/02 15:04:05")
	outModel.Lokasi 		= c.GetString("lokasi")
	outModel.KodeBarang 	= c.GetString("kode")
	outModel.NamaBarang 	= c.GetString("nama")
	outModel.Satuan 		= c.GetString("satuan")
	
	i, _ := strconv.Atoi(c.GetString("jumlah"))
	outModel.Jumlah = i
	fmt.Println(outModel)
	result, err := goodsRepo.OutProccess(outModel)

	if err != nil {
		fmt.Println(result)
		c.Data["messageSuccess"]= "Gagal menambahkan barang"
		c.Layout 				= "template.html"
		c.TplName 				= "incoming.html"
	}
	
	fmt.Println(result)
	c.Data["messageSuccess"] = "Berhasil menambahkan barang"
	c.Redirect("/incoming-goods",302)
}
