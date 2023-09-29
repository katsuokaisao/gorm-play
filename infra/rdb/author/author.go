package author

import (
	"github.com/katsuokaisao/gorm/domain/author"
	"github.com/katsuokaisao/gorm/infra/rdb"
)

type authorRepository struct {
	db rdb.RDB
}

func NewAuthorRepository(db rdb.RDB) author.AuthorRepository {
	return &authorRepository{
		db: db,
	}
}
