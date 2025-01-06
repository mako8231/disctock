package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID      User
	Description string
}

func CreateTask(db *gorm.DB, userID int, description string) {
	var user User
	db.First(&user, userID)
	db.Create(&Task{UserID: user, Description: description})
}
