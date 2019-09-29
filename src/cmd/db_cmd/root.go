package db_cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "db",
	Short: "handle database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This command is to handle database.")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		createDBCmd(),
		createMigrationCmd(),
		upMigrationsCmd(),
		downMigrationsCmd(),
	)
}
