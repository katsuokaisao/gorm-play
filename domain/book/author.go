package book

import (
	"time"
)

type Author struct {
	ID int64 `gorm:"primary_key"`

	Name string `gorm:"name"`

	Books []Book `gorm:"many2many:book_authors"`

	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at"`
}

type AuthorRepository interface {
	FindAll() ([]Author, error)
	FindByID(id int64) (*Author, error)
	Create(author *Author) error
	Update(id int64, name *string) error
	Delete(id int64) error
	DeleteByIDs(ids []int64) error
}
