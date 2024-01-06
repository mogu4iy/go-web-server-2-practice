package core

import (
	dbmodels "go-web-server-2-practice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	DSN string
	Engine *gorm.DB
}

func (db *DB) Init () (err error){
	db.Engine, err = gorm.Open(mysql.New(mysql.Config{
		DSN: db.DSN,
	}), &gorm.Config{})
	if err !=nil {
		return
	}
	sqlDB, err := db.Engine.DB()
	if err !=nil {
		return
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(2)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.Engine.AutoMigrate(&dbmodels.User{})
	if err != nil {
		return
	}
	return
}

