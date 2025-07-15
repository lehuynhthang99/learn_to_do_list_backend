package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	charset_value    string = "utf8mb4"
	parse_time_value string = "True"
	time_value       string = "UTC"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbConfig := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
		charset_value,
		parse_time_value,
		time_value,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("Connected to MySQL")
}
