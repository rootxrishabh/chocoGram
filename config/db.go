package config

import (
	  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Use gorm to connect to mysql server and return the connection
func GetDB() *gorm.DB{
	d, err := gorm.Open("mysql", "root:password@/socialNetwork?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	return d
}