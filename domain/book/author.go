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
}

type AuthorRepository interface {
	FindAll() ([]Author, error)
	FindByID(id int64) (*Author, error)
	FindByName(name string) (*Author, error)
	Create(author *Author) (*Author, error)
	Update(id int64, name *string) error
	Delete(id int64) error
	DeleteByIDs(ids []int64) error
}
