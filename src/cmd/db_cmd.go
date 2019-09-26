package main

import (
	"fmt"
	"github.com/sashiyama/syncnow_server/cmd/db_cmd"
	"os"
)

func main() {
	if err := db_cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
