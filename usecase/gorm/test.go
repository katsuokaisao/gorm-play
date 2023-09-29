package gorm

import (
	"fmt"

	"github.com/katsuokaisao/gorm/domain/book"
)

type GormUseCase interface {
	Test()
}

type gormUseCase struct {
	bookRepository        book.BookRepository
	bookRewViewRepository book.BookReviewRepository
	authorRepository      book.AuthorRepository
}

func NewGormUseCase(
	bookRepository book.BookRepository,
	bookRewViewRepository book.BookReviewRepository,
	authorRepository book.AuthorRepository,
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

	fmt.Println("FindByID(1) result book1-title", b.Title)

	books, err := u.bookRepository.FindAll()
	if err != nil {
		panic(err)
	}

	for i, book := range books {
		fmt.Printf("FindAll() result book%d-title: %s\n", i, book.Title)
		for j, author := range book.Authors {
			fmt.Printf("FindAll() result book%d-author%d-name: %s\n", i, j, author.Name)
		}
		for j, review := range book.BookReviews {
			fmt.Printf("FindAll() result book%d-review%d-review: %s\n", i, j, review.Review)
		}
	}
}
