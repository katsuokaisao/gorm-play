package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/katsuokaisao/gorm/infra/rdb"
	"github.com/katsuokaisao/gorm/infra/rdb/book"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var authorFindByIdCmd = &cobra.Command{
	Use:   "author-find-by-id",
	Short: "Find author by id",
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
		author, err := authorRepo.FindByID(1)
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(author, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

		author, err = authorRepo.FindByID(10000)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("Author not found")
				return
			}
		}
		b, err = json.MarshalIndent(author, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	},
}
