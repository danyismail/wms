package helper

import(
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	RepoReporting "wms/repo/reporting"
	"wms/models"
)

type HelperModule struct{}

var Xlsx *excelize.File

func (hm *HelperModule) ExportToExcell() string{
	repo := RepoReporting.ReportRepo{}
	result, err := repo.All()
	if err != nil {
		fmt.Println("Error saat mengambil data reporting...")
	}
	Xlsx = excelize.NewFile()

	sheet1Name := "Sheet1"
	Xlsx.SetSheetName(Xlsx.GetSheetName(1), sheet1Name)

	//prepare column
	Xlsx.SetCellValue(sheet1Name, "A1", "Nama Barang")
	Xlsx.SetCellValue(sheet1Name, "B1", "Lokasi Gudang")
	Xlsx.SetCellValue(sheet1Name, "C1", "Kode Satuan")
	Xlsx.SetCellValue(sheet1Name, "D1", "Jumlah Barang Masuk")
	Xlsx.SetCellValue(sheet1Name, "E1", "Jumlah Barang Keluar")
	Xlsx.SetCellValue(sheet1Name, "F1", "Sisa Barang")

	errFilterXLS := Xlsx.AutoFilter(sheet1Name, "A1", "F1", "")
	if errFilterXLS != nil {
	    fmt.Println(errFilterXLS)
	}


	y := make([]interface{}, len(result))
	for i, each := range result {
    	y[i] = each
	}
	

	for i, eachInterface := range y {
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), eachInterface.(models.Reporting).NamaBarang)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), eachInterface.(models.Reporting).Lokasi)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), eachInterface.(models.Reporting).Kode)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), eachInterface.(models.Reporting).JumlahMasuk)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), eachInterface.(models.Reporting).JumlahKeluar)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+2), eachInterface.(models.Reporting).Sisa)
	}

	errCreateXLS := Xlsx.SaveAs("./static/laporan-keluar-masuk-barang.xlsx")
	if errCreateXLS != nil {
    	fmt.Println(err)
	}

	/*
	var write http.ResponseWriter
	var read *http.Request
	DownloadExcell(write, read, Xlsx)
	*/

	fmt.Println("Created xlsx success")
	return "Created xls sucess..."
}

/*
func DownloadExcell(w http.ResponseWriter, r *http.Request, file *excelize.File){
	w.Header().Set("Content-Type", "application/octet-stream")
  	w.Header().Set("Content-Disposition", "attachment; filename,=userInputData.xlsx")
  	w.Header().Set("File-Name", "Data.xlsx")
  	w.Header().Set("Content-Transfer-Encoding", "binary")
  	w.Header().Set("Expires", "0")
	errWrite := file.Write(w)
	if errWrite != nil {
		fmt.Println("error download excell")
	}
}
*/

