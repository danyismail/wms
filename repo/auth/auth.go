package repo

import (
	"wms/database"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"wms/models"
)

type AuthRepo struct{
	beego.Controller
}



func(repo *AuthRepo) Login(form models.LoginForm) bool{
	//Logicnya cari di table users yg emailnya == input email, kalo ada 
	//maka ambil passwordnya lalu di compare
	userModel := models.LoginResponse{}
	database.DB.Raw("SELECT username, email, password, role FROM users WHERE email = ?", form.Email).Scan(&userModel)
	match := repo.CheckPassword(userModel.Password, form.Password)
	if !match{
		return false
	}
	return true
}

func (repo *AuthRepo) CheckPassword(hashPwd string, plainPwd string) bool{
	/*
		PENJELASAN
		Fungsi ini membandingkan antara password hash yang ada di db dengan inputan user, returnnya adalah false or true.
		Fungsi ini menerima 2 argumen yaitu string password yg sudah di hash, dan string password dari form input
	*/

	//Kedua argumen dirubah jadi byte code
	byteHash := []byte(hashPwd)
	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}