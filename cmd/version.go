package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cliVersion = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Secrethor CLI version:", cliVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// go build -ldflags "-X 'secrethor-cli/cmd.cliVersion=v0.1.0'" -o secrethor-cli main.go
