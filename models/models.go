package models

// Book represents the structure of a book
type Book struct {
	ID          uint      `gorm:"primarykey"`
	Title       string    `json:"title"`
	AuthorID    uint      `json:"author_id"`
	PublisherID uint      `json:"publisher_id"`
	Price       float64   `json:"price"`
	Author      Author    `gorm:"foreignKey:AuthorID"`
	Publisher   Publisher `gorm:"foreignKey:PublisherID"`
}

// Author represents the structure of an author
type Author struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
}

// Publisher represents the structure of a publisher
type Publisher struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`
}
