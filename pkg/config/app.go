package config

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	
)

var (
	db *gorm.DB
)

func Connect() {
	d, err:=gorm.Open("mysql", "root:vogo1005@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	// this is how mysql is connected to the go project
	// where username:root 
	// password: vogo1005
	// host: localhost:3306
	// databasename: simplerest 
	
	if err != nil{
		panic(err)
	}
	db = d
	fmt.Println("Is connected to Database")
}

func GetDB() *gorm.DB{
	return db
}
