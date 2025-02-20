package cmd

import (
	"fmt"
	"log"
	"myapp/delivery/api"
	"myapp/global"
	"myapp/manager"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cms",
	Short: "Start the CMD http server",
	Run: func(cmd *cobra.Command, args []string) {
		container := manager.NewContainer()
		defer func() {
			if err := container.Close(); err != nil {
				panic(err)
			}
		}()

		router := api.NewRouter(container)

		addr := fmt.Sprintf(":%d", global.GetConfig().Port)

		if err := router.Run(addr); err != nil {
			panic(err)
		}
	},
}

func ExecuteRoot() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
