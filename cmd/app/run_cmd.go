package app

import (
	"fmt"
	"log"

	"github.com/leighlin0511/grpc_template/internal/app"
	"github.com/leighlin0511/grpc_template/internal/app/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var conf config.Configuration

func init() {
	addIntFlag("grpcport", 50051, "gRPC port")
	addIntFlag("httpport", 8080, "HTTP port")
}

func addIntFlag(name string, value int, usage string) {
	runCmd.PersistentFlags().IntP(name, "", value, usage)
	bindPFlags(name)
}

func bindPFlags(name string) {
	if err := viper.BindPFlag(name, runCmd.PersistentFlags().Lookup(name)); err != nil {
		log.Fatalln(fmt.Errorf("unable to bind flag %v: %w", name, err))
	}
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the template server",
	Long:  "run the sample grpc http template server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&conf); err != nil {
			log.Fatal("Unmashal configuration file error:", err)
		}
		app.Run(&conf)
	},
}
