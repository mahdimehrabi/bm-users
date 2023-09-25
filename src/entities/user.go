package entities

type User struct {
	ID        int64
	email     string
	fullname  string
	lastLogin int64    //this field is timestamp and must fill using authorization service and kafka
	lastIPs   []string //user's 10 last logged in ip addresses and must fill using authorization service and kafka
}
