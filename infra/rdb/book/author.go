package book

import (
	"github.com/katsuokaisao/gorm/domain/book"
	"github.com/katsuokaisao/gorm/infra/rdb"
)

type authorRepository struct {
	db rdb.RDB
}

func NewAuthorRepository(db rdb.RDB) book.AuthorRepository {
	return &authorRepository{
		db: db,
	}
}
