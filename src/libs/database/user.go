package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"libs/constant"
)

type User struct {
	User_Id   	uint64
	Password 	string
	User_type	string
	Name    	string

}



func UserDeleteByName(name string) {
	db, err := gorm.Open("mysql", constant.HisDbConnectStr)
	db.SingularTable(true)
	if err != nil {
		fmt.Println("connect mysql error", err)
		return
	}
	fmt.Println("Connection done")
	defer db.Close()

	db.Where("name = ?", name).Delete(&User{})
	fmt.Println("Delete user:", name)
}

func UserFindByName(name string) bool {
	db, err := gorm.Open("mysql", constant.HisDbConnectStr)
	db.SingularTable(true)
	if err != nil {
		fmt.Println("connect mysql error", err)
		return false
	}
	fmt.Println("Connection done")
	defer db.Close()

	var count int = 0
	db.Where("name = ?", name).Find(&User{}).Count(&count)

	fmt.Println("Find user:", count)
	if count >0 {
		return true
	} else {
		return false
	}
}
