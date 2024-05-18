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

var authorDeleteCmd = &cobra.Command{
	Use:   "author-delete",
	Short: "Delete author",
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
		var (
			name   string = "John Doe"
			author *dbook.Author
			err    error
		)
		author, err = authorRepo.FindByName(name)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
			author, err = authorRepo.Create(&dbook.Author{
				Name: name,
			})
			if err != nil {
				panic(err)
			}
		}
		b, err := json.MarshalIndent(author, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		fmt.Println("Deleting author...")

		if err := authorRepo.Delete(author.ID); err != nil {
			panic(err)
		}

		if err := authorRepo.Delete(author.ID); err != nil {
			panic(err)
		}
	},
}
