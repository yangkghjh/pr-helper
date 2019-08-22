package server

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yangkghjh/pr-helper/server/apis"
)

var configFilePath string

func init() {
	Run.Flags().StringVarP(&configFilePath, "config", "c", "", "config file path")
}

// Run start api server
var Run = &cobra.Command{
	Use:   "run",
	Short: "run the api server",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	// load config file
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// start server
	apis.Run()
}
