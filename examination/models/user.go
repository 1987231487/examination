package models

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Level int `json:"level"`
}

//func init(){
//	db.Table("user").CreateTable(&User{})
//}
func CheckUser(username, password string) bool {
	var user User
	db.Table("user").Select("id").Where(User{Username : username, Password : password}).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}
func CheckName(username string) bool{
	var user User
	db.Table("user").Where("username=?",username).Find(&user)
	if user.ID>0 {
		return true
	}
	return false
}
func CreateUser(username, password,email string,level int)  {
	user:=User{Username : username, Password : password,Level: level,Email: email}
	db.Table("user").Create(&user)
}
func GetUserPower(UserName string)(int){
	var user User
	db.Table("user").Where("username=?",UserName).First(&user)
	return user.Level
}
