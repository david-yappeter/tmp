//go:build devtools

package cmd

import (
	"fmt"
	"myapp/database/seeder"
	"myapp/database/seeder/production_seeder"
	"myapp/global"
	"myapp/manager"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newSeedCommand())
}

func newSeedCommand() *cobra.Command {
	var flagProduction bool

	cmd := &cobra.Command{
		Use:   "seed",
		Short: "Seed the database with records",
		Long:  "Seed the database with records",
		Run: func(_ *cobra.Command, args []string) {
			global.DisableDebug()

			container := manager.NewContainer()
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			repositoryManager := container.RepositoryManager()

			if len(args) > 0 {
				fmt.Printf("Seeder for table `%s` not found\n", args[0])
				return
			}

			if flagProduction {
				production_seeder.SeedAll(repositoryManager)
			} else {
				seeder.SeedAll(repositoryManager)
			}
		},
	}

	cmd.Flags().BoolVarP(&flagProduction, "production", "p", false, "seed using production data")

	for tableName := range seeder.Seeders {
		cmd.AddCommand(newSeedTableCommand(tableName))
	}

	return cmd
}

func newSeedTableCommand(tableName string) *cobra.Command {
	var flagProduction bool

	cmd := &cobra.Command{
		Use:   tableName,
		Short: fmt.Sprintf("Seed table %s with records", tableName),
		Long:  fmt.Sprintf("Seed table %s with records", tableName),
		Run: func(_ *cobra.Command, _ []string) {
			global.DisableDebug()

			container := manager.NewContainer()
			defer func() {
				if err := container.Close(); err != nil {
					panic(err)
				}
			}()

			repositoryManager := container.RepositoryManager()

			if flagProduction {
				production_seeder.Seed(repositoryManager, tableName)
			} else {
				seeder.Seed(repositoryManager, tableName)
			}
		},
	}

	if _, exist := production_seeder.Seeders[tableName]; exist {
		cmd.Flags().BoolVarP(&flagProduction, "production", "p", false, "seed using production data")
	}

	return cmd
}
