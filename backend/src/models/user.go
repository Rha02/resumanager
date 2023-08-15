package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
}
