package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// create a seed file for book model
func Seed(db *gorm.DB) {
	books := []Book{
		{
			Title:       "The Alchemist",
			Author:      "Paulo Coelho",
			Description: "A novel about a young Andalusian shepherd named Santiago who travels from his homeland in Spain to the Egyptian desert in search of a treasure buried in the Pyramids.",
			Price:       10.50,
		},
		{
			Title: "The Little Prince",
			Author: "Antoine de Saint-Exupéry",
			Description: "A poetic tale self-illustrated in watercolours in which a pilot stranded in the desert meets a young prince fallen to Earth from a tiny asteroid.",
			Price: 12.50,
		},
		{
			Title: "The Catcher in the Rye",
			Author: "J.D. Salinger",
			Description: "A novel about a young",
			Price: 15.50,
		},
	}

	// loop through the books and create them
	for _, book := range books {
		result := db.Create(&book)
		fmt.Print(result.Error)
		if result.Error!=nil {
			log.Fatalf("Error seeding books: %v", err)
		}
	}

}
