package models

type Account struct {
	ID      int     `json:"id"`
	UserID  int     `json:"userId"`
	Balance float32 `json:"balance"`
	Created int     `json:"created"`
	Updated int     `json:"updated"`
	Active  uint8   `json:"completed"`
}

type Deposit struct {
	Amount float32 `json:"amount"`
}

type PaymentParams struct {
	UserID    int     `gorm:"index;type:int" json:"userId"`
	Amount    float32 `gorm:"type:float" json:"amount"`
	RequestID string  `gorm:"index;type:string" json:"requestId"`
}
