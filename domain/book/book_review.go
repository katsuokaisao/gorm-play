package book

import "time"

type BookReview struct {
	ID     int64 `gorm:"primary_key"`
	BookID int64 `gorm:"book_id"`

	Rating  int    `gorm:"rating"`
	RewView string `gorm:"review"`

	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at"`
}

type BookReviewRepository interface {
}
