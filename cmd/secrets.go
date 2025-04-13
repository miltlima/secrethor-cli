package cmd

import (
	"github.com/spf13/cobra"
)

// secretsCmd groups subcommands related to Kubernetes secrets.
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Manage and inspect Kubernetes Secrets",
	Long:  "Subcommands to scan, search, and manage Kubernetes Secrets in your cluster.",
}

func init() {
	rootCmd.AddCommand(secretsCmd)
}
