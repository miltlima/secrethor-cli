package cmd

import (
	"github.com/miltlima/secrethor-cli/internal/secrethor"
	"github.com/spf13/cobra"
)

var (
	namespace string
	output    string
	verbose   bool
)

var orphanCmd = &cobra.Command{
	Use:   "orphan",
	Short: "Scan and report orphaned Kubernetes Secrets",
	RunE: func(cmd *cobra.Command, args []string) error {
		return secrethor.Check(namespace, output, verbose)
	},
}

func init() {
	orphanCmd.Flags().StringVarP(&namespace, "namespace", "n", "all", "Namespace to scan (or 'all')")
	orphanCmd.Flags().StringVarP(&output, "output", "o", "table", "Output format: table | json | yaml")
	orphanCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose debug output")
	secretsCmd.AddCommand(orphanCmd)
}
