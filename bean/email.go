package bean

import ()

type Email struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Mailtype string `json:"type"`
}

type EmailExtra struct {
	Email    Email             `json:"email"`
	TempName string            `json:"tempname"`
	Parse    map[string]string `json:"parse"`
}
