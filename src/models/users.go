package models

type User struct {
	ID       int
	Login    string
	Password string
	Created  int
	Updated  int
	Active   int
}
