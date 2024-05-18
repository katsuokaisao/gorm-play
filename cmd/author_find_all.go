package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/katsuokaisao/gorm/infra/rdb"
	"github.com/katsuokaisao/gorm/infra/rdb/book"
	"github.com/spf13/cobra"
)

var authorFindAll = &cobra.Command{
	Use:   "author-find-all",
	Short: "Find all authors",
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
		authors, err := authorRepo.FindAll()
		if err != nil {
			panic(err)
		}
		if len(authors) == 0 {
			fmt.Println("No authors found")
			return
		}
		b, err := json.MarshalIndent(authors, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	},
}
