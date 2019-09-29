package db_cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func createDBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create_db",
		Short: "create database",
		RunE: func(cmd *cobra.Command, args []string) error {
			out, error := exec.Command("psql", "-h", "db", "-U", "postgres", "-w", "-c", "create database syncnow_development;").CombinedOutput()
			if error != nil {
				fmt.Println("Command Exec Error.")
			}
			fmt.Printf("result: \n%s", string(out))

			return nil
		},
	}

	return cmd
}
