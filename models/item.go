package models

type Item struct {

	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name"`

	Price string `json:"price"`

	Qty int `json:"qty"`
}
