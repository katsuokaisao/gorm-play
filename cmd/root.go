package cmd

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/katsuokaisao/gorm/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "",
}

var cfg *config.Config

func init() {
	rootCmd.AddCommand(authorFindAll)
	rootCmd.AddCommand(authorFindByIdCmd)
	rootCmd.AddCommand(authorCreateCmd)
	rootCmd.AddCommand(authorDeleteCmd)
	initConfig()
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	c := config.Config{
		DB: &config.DB{},
	}
	if err := env.Parse(&c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	cfg = &c
}

func Execute() {
	rootCmd.Execute()
}
