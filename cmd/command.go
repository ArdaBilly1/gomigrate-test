package cmd

import (
	"fmt"
	"gomigrate-test/migration"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(migration.Execute())
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("cmd not run")
			os.Exit(0)
		}
	},
}
