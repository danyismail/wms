package api

import (
	"github.com/astaxie/beego"
	Reporting "wms/repo/reporting"
	"fmt"
	WHelper"wms/helper"
)


type ReportingController struct {
	beego.Controller
}

func (c *ReportingController) All() {
	repo := Reporting.ReportRepo{}
	result, err := repo.All()
	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "reporting/in-out.html" //buat load halaman
	}
	fmt.Println(result)
	c.Data["list"] = result
	c.Layout = "template.html"
	c.TplName = "reporting/in-out.html" //buat load halaman
}

func (c *ReportingController) Export() {
	exp := WHelper.HelperModule{}
	result := exp.ExportToExcell()
	c.Data["excellSuccess"] = result
	c.Layout = "template.html"
	c.TplName = "reporting/success-export.html" //buat load halaman
}



