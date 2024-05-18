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

var authorUpdateCmd = &cobra.Command{
	Use:   "author-update",
	Short: "Update author",
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
		fmt.Println("Updating author...")

		author.Name = "Jane Doe"
		if err := authorRepo.Update(author.ID, &author.Name); err != nil {
			panic(err)
		}

		author, err = authorRepo.FindByID(author.ID)
		if err != nil {
			panic(err)
		}
		b, err = json.MarshalIndent(author, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

		if err := authorRepo.Delete(author.ID); err != nil {
			panic(err)
		}
	},
}
