package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	Fullname  string
	LastLogin int64  //this field is timestamp and must fill using authorization service and kafka
	LastIPs   string //user's 10 last logged in ip addresses and must fill using authorization service and kafka (seperated by ",")
	Password  string //user password
}
