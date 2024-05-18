package cmd

import (
	"fmt"

	dbook "github.com/katsuokaisao/gorm/domain/book"
	"github.com/katsuokaisao/gorm/infra/rdb"
	"github.com/katsuokaisao/gorm/infra/rdb/book"
	"github.com/spf13/cobra"
)

var authorDeleteIdsCmd = &cobra.Command{
	Use:   "author-delete-ids",
	Short: "Delete authors by ids",
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

		authors := []dbook.Author{
			{
				Name: "Author 1",
			},
			{
				Name: "Author 2",
			},
		}
		if err := authorRepo.CreateBulk(authors); err != nil {
			panic(err)
		}

		ids := []int64{}
		for _, author := range authors {
			ids = append(ids, author.ID)
		}
		fmt.Printf("ids: %+v\n", ids)
		if err := authorRepo.DeleteByIDs(ids); err != nil {
			panic(err)
		}

		if err := authorRepo.DeleteByIDs(ids); err != nil {
			panic(err)
		}
	},
}
