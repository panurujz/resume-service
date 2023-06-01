package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Open() (db *gorm.DB) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=P@ssw0rd dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "goapp.",
		},
	})

	if err != nil {
		log.Fatal("Fail to connect database.")
	}

	// db.AutoMigrate(&models.User{})

	fmt.Println("Connect database success.")

	return db
}
