package models

import (
"github.com/jinzhu/gorm"
)

var pool *gorm.DB

func GetDB() *gorm.DB {
	if pool == nil {
		db := MakeDB()
		pool=db
	}

	return pool
}

func MakeDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return db
}
