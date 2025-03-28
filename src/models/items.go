package models

type Item struct {
	ID             int     `gorm:"index;type:int" json:"id"`
	Title          string  `gorm:"type:text" json:"title"`
	Description    string  `gorm:"type:text" json:"description"`
	Price          float32 `gorm:"type:float" json:"price"`
	CategoryID     int     `gorm:"index;type:int" json:"category_id"`
	Created        string  `gorm:"type:text" json:"created"`
	Updated        string  `gorm:"type:text" json:"updated"`
	UserCategoryID int     `gorm:"index;type:int" json:"user_category_id"`
	Active         uint8   `gorm:"index;type:int" json:"active"`
}

func (s Item) TableName() string { return "items" }

type BuyItem struct {
	ID     int `json:"id"`
	Volume int `json:"volume"`
}

type ItemCategory struct {
	ID             int    `gorm:"index;type:int" json:"id"`
	Title          string `gorm:"type:text" json:"title"`
	Description    string `gorm:"type:text" json:"description"`
	Created        string `gorm:"type:text" json:"created"`
	Updated        string `gorm:"type:text" json:"updated"`
	UserCategoryID int    `gorm:"index;type:int" json:"user_category_id"`
	Active         uint8  `gorm:"index;type:int" json:"active"`
}

func (s ItemCategory) TableName() string { return "item_category" }
