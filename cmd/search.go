package cmd

import (
	"github.com/miltlima/secrethor-cli/internal/secrethor"
	"github.com/spf13/cobra"
)

var searchNamespace string

var searchCmd = &cobra.Command{
	Use:   "search [SECRET_NAME]",
	Short: "Search for a specific Kubernetes Secret",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		secretName := args[0]
		return secrethor.SearchSecret(secretName, searchNamespace)
	},
}

func init() {
	searchCmd.Flags().StringVarP(&searchNamespace, "namespace", "n", "all", "Namespace to search (or 'all')")
	secretsCmd.AddCommand(searchCmd)
}
