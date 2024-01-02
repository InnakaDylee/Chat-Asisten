package model

type Message struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Role 	  string `json:"role"`
	Content   string `json:"content"`
}