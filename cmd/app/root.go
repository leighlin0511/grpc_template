package app

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var configPath = "./"
var configName = "config"

var rootCmd = &cobra.Command{
	Use: "app",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "", "", "config file")
	if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		log.Fatalln(fmt.Errorf("unable to initialize flag config: %w", err))
	}
}

func initConfig() {
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(configPath)
		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println("ReadInConfig:", err)
	}
}

func Execute() {
	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
