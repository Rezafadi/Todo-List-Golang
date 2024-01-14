package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=technical-test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic("failed to connect database")
	}


	return db
}

