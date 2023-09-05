package models

import (
	"task-5-pbi-btpns-arthagusfiputra/app"
	"github.com/google/uuid"
	"html"
	"strings"
)

type User struct{
	// use gorm for the models
	ID string `gorm:"primary_key; unique" json:"id"`
	Username string `gorm:"size:255;not null;" json:"username"`
	Email string `gorm:"size:255;not null;unique"  json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	Photos Photo `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"photos"`
}

type Photo struct{
	ID       int       `gorm:"primary_key;auto_increment" json:"id"`
	Title    string    `gorm:"size:255;not null" json:"title"`
	Caption  string    `gorm:"size:255;not null" json:"caption"`
	PhotoUrl string    `gorm:"size:255;not null;" json:"photo_url"`
	UserID   string    `gorm:"not null" json:"user_id"`
	Owner    app.Owner `gorm:"owner"`
}
func (u *User) Init() {
	u.ID = uuid.New().String() //generate uuid                             
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}