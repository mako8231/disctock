package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string
	UserToken   string
	CreatedDate time.Time
}

func CreateUser(db *gorm.DB, userName string) User {
	var user User
	hash := md5.Sum([]byte(userName))
	token := hex.EncodeToString(hash[:])
	user = User{UserName: userName, UserToken: token, CreatedDate: time.Now()}

	db.Create(&user)
	return user
}

func ShowUser(db *gorm.DB, userID int) User {
	var user User
	db.First(&user, userID)
	return user
}
