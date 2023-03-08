package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {

	var err error

	connection_string := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Error connecting to Db", err)
	}
}
