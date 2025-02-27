package models

type NotificationMessage struct {
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
