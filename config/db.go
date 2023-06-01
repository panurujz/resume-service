package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Open() (db *gorm.DB) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection, _ := sql.Open("pgx", os.Getenv("DB_URL"))

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn:                 connection,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: os.Getenv("DB_SCHEMA"),
		},
	})

	if err != nil {
		log.Fatal("Fail to connect database.")
	}

	// db.AutoMigrate(&models.User{})

	return db
}
