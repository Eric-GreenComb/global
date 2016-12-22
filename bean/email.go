package bean

import ()

// Email email struct
type Email struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Mailtype string `json:"type"`
}

// EmailExtra email extra struct
type EmailExtra struct {
	Email    Email             `json:"email"`
	TempName string            `json:"tempName"`
	Parse    map[string]string `json:"parse"`
}
