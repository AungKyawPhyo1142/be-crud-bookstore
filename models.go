package main

type Book struct {
	ID	 uint `gorm:"primaryKey"`
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}