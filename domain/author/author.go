package author

import "time"

type Author struct {
	ID int64 `gorm:"primary_key"`

	Name string `gorm:"name"`

	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at"`
}

type AuthorRepository interface {
}
