package controller

import (
	"encoding/json"
	_ "fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MODEL User
type User struct{
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
}

//membuat table User secara otomatis
func UserMigration(){
	db, err := gorm.Open(sqlite.Open("db.sqlite"))

	if err != nil {
		panic("gagal membuat table user")
	}

	// migrate user table
	db.AutoMigrate(&User{})
}

// Controller user_create
func User_create(res http.ResponseWriter, req *http.Request){
	//membuka connection 
	db, err := gorm.Open(sqlite.Open("db.sqlite"))
	if err != nil{
		panic("gagal membuka database")
	}

	var user User

	json.NewDecoder(req.Body).Decode(&user)

	query := db.Create(&user)

	if query.Error != nil{
		
		type errorMessage struct{
			Success bool
			Msg interface{}
		}
		
		result, _ := json.Marshal(&errorMessage{
			Success : false,
			Msg: query.Error,
		})
		
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(result)
		return
	}

	result, _ := json.Marshal(&user)

	//jika semuanya berhasil
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(result)


}
