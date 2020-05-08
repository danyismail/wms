package repo

import(
	"wms/database"
	"wms/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct{
	beego.Controller
}

func (repo *UsersRepo) All() (result []models.Users ,err error){

    var users []models.Users
	if result := database.DB.Find(&users); result.Error != nil {
		return users, result.Error
	}

	return users, nil
}	

func (repo *UsersRepo) Add(form models.Users) (string, error){
	hashPassword := repo.HashingWithSalt(form.Password)
	form.Password = hashPassword
	database.DB.Create(&form)
	return "berhasil menambahkan users", nil
}

func (repo *UsersRepo) HashingWithSalt(password string) string{
	//rubah password ke byte, karena syarat untuk hashing data nya harus dalam bentuk byte
	bytePassword := []byte(password)

	//Lakukan hashing dan salt password . Note : returnnya byte
	hashAndSalt, errHash := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if errHash != nil {
		return "proses hashing error"
	}

	//casting dari byte ke string dan return hasilnya
	return string(hashAndSalt)
}







