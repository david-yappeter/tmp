//go:build devtools

package cmd

import (
	"myapp/global"
	"myapp/manager"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newMigrateFreshCommand())
}

func newMigrateFreshCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-fresh",
		Short: "Remove and migrate the database table",
		Long:  "Remove and migrate the database table",
		Run: func(_ *cobra.Command, _ []string) {
			global.DisableDebug()

			container := manager.NewContainer()
			if err := container.RefreshDB(); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
