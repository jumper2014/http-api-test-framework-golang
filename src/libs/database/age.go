package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Age struct {
	Id   int
	Name string
}

func FindName(db *gorm.DB, name string) bool {
	var age Age
	db.First(&age, "name = ?", name)
	fmt.Printf("%v\n", age)
	if age.Name != name {
		fmt.Printf("not found %v", age.Name)
		return false
	} else {
		return true
	}

}

func DeleteName(db *gorm.DB, name string) {
	db.Where("name = ?", name).Delete(&Age{})
}