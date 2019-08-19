package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB

func DatabaseInit(uri string) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	//
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//
	if err := db.HasTable(&User{}); !err {
		db.CreateTable(&User{})
	}
	DB = db
}
