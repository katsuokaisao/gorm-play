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

func (r *authorRepository) FindAll() ([]book.Author, error) {
	var authors []book.Author
	if err := r.db.NewSession().Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *authorRepository) FindByID(id int64) (*book.Author, error) {
	var author book.Author
	if err := r.db.NewSession().Find(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *authorRepository) Create(author *book.Author) error {
	return r.db.NewSession().Create(author).Error
}

func (r *authorRepository) Update(id int64, name *string) error {
	input := map[string]interface{}{}
	if name != nil {
		input["name"] = *name
	}
	return r.db.NewSession().Model(&book.Author{ID: id}).Updates(input).Error
}

func (r *authorRepository) Delete(id int64) error {
	return r.db.NewSession().Delete(&book.Author{}, id).Error
}

func (r *authorRepository) DeleteByIDs(ids []int64) error {
	return r.db.NewSession().Delete(&book.Author{}, ids).Error
}

