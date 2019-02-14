package db

import (
"github.com/jinzhu/gorm"
)

var pool *gorm.DB

func GetDB() (*gorm.DB, error) {
	if pool == nil {
		db,err:=MakeDB()
		if err != nil {
			return nil,err
		}
		pool=db
	}

	return pool, nil
}

func MakeDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	return db, err
}
