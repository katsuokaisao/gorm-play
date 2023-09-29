package book

import (
	"github.com/katsuokaisao/gorm/domain/book"
	"github.com/katsuokaisao/gorm/infra/rdb"
)

type bookReviewRepository struct {
	db rdb.RDB
}

func NewBookReviewRepository(db rdb.RDB) book.BookReviewRepository {
	return &bookReviewRepository{
		db: db,
	}
}
