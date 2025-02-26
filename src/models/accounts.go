package models

type Account struct {
	ID      int     `json:"id"`
	UserID  int     `json:"userId"`
	Balance float32 `json:"price"`
	Created int     `json:"created"`
	Updated int     `json:"updated"`
	Active  uint8   `json:"completed"`
}
