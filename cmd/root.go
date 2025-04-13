package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secrethor-cli",
	Short: "Secrethor CLI - Manage and validate Kubernetes Secrets",
	Long:  "Secrethor is a CLI tool for scanning, validating, and managing Kubernetes Secrets across workloads.",
}

// Execute runs the CLI root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
