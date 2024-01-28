package database

import (
	"fmt"
	"github.com/BenzoFuryWolf/MyProject/config"
	"github.com/BenzoFuryWolf/MyProject/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	HOST := config.Config("DB_HOST")
	PORT := config.Config("DB_PORT")
	USER := config.Config("DB_USER")
	PASSWRD := config.Config("DB_PASSWORD")
	DBNAME := config.Config("DB_NAME")
	SSLMODE := config.Config("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWRD, DBNAME, SSLMODE)

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("faild to connect to DB")
	}
	fmt.Println("Connection Opened to DB")

	DB.AutoMigrate(&model.Person_info{})
	fmt.Println("Database Migrated")
}
