package models

type Item struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Created     int     `json:"created"`
	Updated     int     `json:"updated"`
	Active      uint8   `json:"active"`
}
