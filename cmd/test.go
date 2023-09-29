package cmd

import (
	"github.com/katsuokaisao/gorm/inject/test"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "gorm test",
	Long:  `gorm test`,
	Run: func(cmd *cobra.Command, args []string) {
		test.SetApplication(&cfg).Run()
	},
}
