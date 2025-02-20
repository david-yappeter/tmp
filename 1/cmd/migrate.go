//go:build tools

package cmd

import (
	"myapp/global"
	"myapp/manager"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newMigrateCommand())
}

func newMigrateCommand() *cobra.Command {
	var (
		isRollingBack bool
		steps         int
		force         *int
	)

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database table",
		Long:  "Migrate the database table",
		Run: func(_ *cobra.Command, _ []string) {
			global.DisableDebug()

			container := manager.NewContainer()
			if err := container.MigrateDB(global.GetMigrationDir(), isRollingBack, steps, force); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&isRollingBack, "rollback", "", false, "Indicate whether migration is rollback or not")
	cmd.Flags().IntVarP(&steps, "steps", "s", 0, "Specify steps if want to migrate n number of migrations")
	cmd.Flags().Var(newOptionalInt(&force), "force", "Specify force flag as an optional parameter")

	return cmd
}

// optionalInt is a custom type that implements the pflag.Value interface
type optionalInt struct {
	value **int
}

func newOptionalInt(target **int) *optionalInt {
	return &optionalInt{value: target}
}

func (o *optionalInt) Set(val string) error {
	v, err := strconv.Atoi(val)
	if err != nil {
		return err
	}
	*o.value = &v
	return nil
}

func (o *optionalInt) Type() string {
	return "int"
}

func (o *optionalInt) String() string {
	if o.value == nil || *o.value == nil {
		return "nil"
	}
	return strconv.Itoa(**o.value)
}
