package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"libs/database"
)



type Product struct {
	Id   int
	Name string
}

func main() {
	db, err := gorm.Open("mysql", database.GormConnectStr)
	db.SingularTable(true)
	if err != nil {
		panic("connect mysql error")
	}
	defer db.Close()

	fmt.Println("Connect done")

	fmt.Println(db.HasTable(&Product{}))
	fmt.Println(db.HasTable("user"))


	product := Product{Name: "mariadb"}

	fmt.Println(db.NewRecord(product)) // => returns `true` as primary key is blank

	fmt.Println(db.Create(&product))

	db.NewRecord(product) // => return `false` after `user` created

	// Create

	fmt.Println("create done")

	db.First(&product)

	database.FindName(db, "dahui")
	database.DeleteName(db, "dahui")
	// Read
	//var user User
	//for i := 0; i<3; i=i+1{
	//	db.First(&user, i) // find user with id 1
	//	fmt.Println(user.Name)
	//}
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	//
	//// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// Delete - delete product
	//db.Delete(&product)
}