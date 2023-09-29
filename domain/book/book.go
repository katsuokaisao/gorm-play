package book

import "time"

type Book struct {
	ID int64 `gorm:"primary_key"`

	Title string `gorm:"title"`
	ISBN  string `gorm:"isbn"`

	BookReviews []BookReview `gorm:"foreignkey:BookID"`

	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at"`
}

type BookRepository interface {
	FindByID(id int) (*Book, error)
	FindAll() ([]Book, error)
}
