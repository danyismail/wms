package api

import (
	"github.com/astaxie/beego"
	"webgudang/repo/goods"
	"webgudang/models"
	"fmt"
	"strconv"
	"math/rand"
	"time"
)


type MainController struct {
	beego.Controller
}


func (c *MainController) All() {
	goodsRepo := repo.GoodsRepo{}

	result, err := goodsRepo.GetAll()

	if err != nil {
		c.Data["pesan"] = "Error"
		c.Layout = "template.html"
		c.TplName = "incoming.html"
	}
	
	c.Data["incoming"] = result
	c.Layout = "template.html"
	c.TplName = "incoming.html"
}

func (c *MainController) Add() {
	goodsRepo := repo.GoodsRepo{}

	goodsModel :=  models.Incoming{}
	
	newTrx := generateTrxID()
	
	t := time.Now()
	
	//mau nampung data dari form
	goodsModel.TransaksiId = newTrx
	goodsModel.Tanggal = t.Format("2006/01/02 15:04:05")
	goodsModel.Lokasi = c.GetString("lokasi")
	goodsModel.KodeBarang = c.GetString("kode")
	goodsModel.NamaBarang = c.GetString("nama")
	goodsModel.Satuan = c.GetString("satuan")
	
	i, _ := strconv.Atoi(c.GetString("jumlah"))
	goodsModel.Jumlah = i


	result, err := goodsRepo.Add(goodsModel)

	if err != nil {
		fmt.Println(result)
		c.Data["messageSuccess"] = "Gagal menambahkan barang"
		c.Layout = "template.html"
		c.TplName = "incoming.html"
	}
	
	fmt.Println(result)
	c.Data["messageSuccess"] = "Berhasil menambahkan barang"
	c.Redirect("/incoming-goods",302)
}

//Incoming Detail
func (c *MainController) Detail() {
	goodsRepo := repo.GoodsRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)
	result, err := goodsRepo.GetById(i)

	if err != nil {
		fmt.Println("Error repo",err)
	}
	
	fmt.Println("error result",result)
	c.Data["detail"] = result
	c.Layout = "template.html"
	c.TplName = "detail.html"
}

func (c *MainController) Delete() {
	goodsRepo := repo.GoodsRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)

	num, err := goodsRepo.Delete(i)

	if err != nil {
		fmt.Println("Error repo",err)
	}
	c.Data["pesan"] = string(num)
	c.Redirect("/incoming-goods",302)
}

//Helper Function
func generateTrxID() string{
	rand.Seed(time.Now().Unix())
    
	pattern := "TRX"
	num1 := randomWithRange(100, 500)
	num2 := randomWithRange(501, 999)

	s1 := strconv.Itoa(num1)
	s2 := strconv.Itoa(num2)

	generate := pattern + "-" + s1 + "-" + s2
	return generate
}


func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}




