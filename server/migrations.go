package server

import (
	"log"

	"github.com/mako8231/disctock/server/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Printf("Running migrations...")
	db.AutoMigrate(&model.Task{})
	log.Printf("Done")
}
