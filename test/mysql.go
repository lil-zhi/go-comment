package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB() (*gorm.DB, error) {
	gormConfig := gorm.Config{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=true&loc=Local", "root", "kzt19981211", "81.70.169.195", 3306, "commentDB", "utf8mb4", "2s")
	db, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.Debug().DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
