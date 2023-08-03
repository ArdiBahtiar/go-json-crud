package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=arjuna.db.elephantsql.com user=ixjuatqq password=P4unoyq0MJKCA5jyHcDhwyjEGCtqG-7r dbname=ixjuatqq port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal nyambung ke Database")
	}
}
