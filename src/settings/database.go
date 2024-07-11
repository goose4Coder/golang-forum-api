package settings

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {

	if gin.Mode() == gin.DebugMode {
		db, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Cannot connect to sqlite database!")
		}
		Database = db
	}

}
