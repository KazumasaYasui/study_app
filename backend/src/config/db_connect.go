package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "studyapp"
	PASS := "studyapp_pw"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "studyapp_development"
	OPTION := "?parseTime=true&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
