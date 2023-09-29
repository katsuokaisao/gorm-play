package gorm

import (
	"fmt"

	"github.com/katsuokaisao/gorm/domain/author"
	"github.com/katsuokaisao/gorm/domain/book"
)

type GormUseCase interface {
	Test()
}

type gormUseCase struct {
	bookRepository        book.BookRepository
	bookRewViewRepository book.BookReviewRepository
	authorRepository      author.AuthorRepository
}

func NewGormUseCase(
	bookRepository book.BookRepository,
	bookRewViewRepository book.BookReviewRepository,
	authorRepository author.AuthorRepository,
) GormUseCase {
	return &gormUseCase{
		bookRepository:        bookRepository,
		bookRewViewRepository: bookRewViewRepository,
		authorRepository:      authorRepository,
	}
}

func (u *gormUseCase) Test() {
	b, err := u.bookRepository.FindByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(*b)
}
