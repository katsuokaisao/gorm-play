package cmd

import (
	"github.com/katsuokaisao/gorm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "",
}

var cfg config.Config

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(testCmd)
}

func initConfig() {
	cfgFileName := "settings/setting.toml"

	viper.SetConfigFile(cfgFileName)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
}

func Execute() {
	rootCmd.Execute()
}
