package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pointSystem/model"
)

func ConnectDataBase() *gorm.DB {
	username := "root"
	password := "root"
	host := "tcp(host.docker.internal:33066)"
	database := "database_user"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.User{}, &model.Point{})

	return db
}
