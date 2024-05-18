package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	dbook "github.com/katsuokaisao/gorm/domain/book"
	"github.com/katsuokaisao/gorm/infra/rdb"
	"github.com/katsuokaisao/gorm/infra/rdb/book"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var authorCreateCmd = &cobra.Command{
	Use:   "author-create",
	Short: "Create author",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("r: %+v\n", cfg.DB)

		r := rdb.NewRDB(
			rdb.Config{
				Driver:   cfg.DB.Driver,
				Address:  cfg.DB.Address,
				Username: cfg.DB.Username,
				Password: cfg.DB.Password,
				Database: cfg.DB.Database,
				Debug:    cfg.DB.Debug,
			},
		)
		authorRepo := book.NewAuthorRepository(r)
		author := &dbook.Author{
			Name: "John Doe",
		}
		a, err := authorRepo.Create(author)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				fmt.Println("Author already exists")
				return
			}
			panic(err)
		}
		b, err := json.MarshalIndent(a, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

		authorDup := &dbook.Author{
			Name: "John Doe",
		}
		ad, err := authorRepo.Create(authorDup)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				fmt.Println("Author already exists")
				return
			}
			panic(err)
		}
		b, err = json.MarshalIndent(ad, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	},
}
