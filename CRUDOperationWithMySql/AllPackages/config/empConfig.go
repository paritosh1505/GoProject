package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbconnection *gorm.DB

func ConnectDB() *gorm.DB{
	connectionName := "root:root@tcp(127.0.0.1:3306)/mygodb?charset=utf8&parseTime=True&loc=Local"
	dbval,err:= gorm.Open(mysql.Open(connectionName),&gorm.Config{})
	if err!=nil{
		fmt.Println("Error while creating DB",err)
	}
	return dbval
}