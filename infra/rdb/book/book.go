package book

import (
	"github.com/katsuokaisao/gorm/domain/book"
	"github.com/katsuokaisao/gorm/infra/rdb"
)

type bookRepository struct {
	db rdb.RDB
}

func NewBookRepository(db rdb.RDB) book.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (repo *bookRepository) FindByID(id int) (*book.Book, error) {
	sess := repo.db.NewSession()

	var book book.Book
	if err := sess.First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
