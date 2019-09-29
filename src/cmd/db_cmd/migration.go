package db_cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func createMigrationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create_migration",
		Short: "create migration",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, error := exec.Command("migrate", "create", "-ext", "sql", "-dir", "db/migrations", "-seq", args[0]).CombinedOutput()
			if error != nil {
				fmt.Printf("Command Exec Error.")
			}
			fmt.Printf("result: \n%s migration file was created.", args[0])

			return nil
		},
	}

	return cmd
}

func upMigrationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up_migration",
		Short: "run migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			out, error := exec.Command("migrate", "-database", os.Getenv("POSTGRESQL_URL"), "-path", "db/migrations", "up").CombinedOutput()
			if error != nil {
				fmt.Println("Command Exec Error.")
			}
			fmt.Printf("result: \n%s", string(out))

			return nil
		},
	}

	return cmd
}

func downMigrationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down_migration",
		Short: "rollback migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			out, error := exec.Command("migrate", "-database", os.Getenv("POSTGRESQL_URL"), "-path", "db/migrations", "down", "-all").CombinedOutput()
			if error != nil {
				fmt.Println("Command Exec Error.")
			}
			fmt.Printf("result: \n%s", string(out))

			return nil
		},
	}

	return cmd
}
