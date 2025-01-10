package models

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
	Active   int    `json:"active"`
}
