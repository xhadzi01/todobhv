package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		errMsg := "Can't connect to db"
		log.Fatal(errMsg)
		panic(errMsg)
	}
}
