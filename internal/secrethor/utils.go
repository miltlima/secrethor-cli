package secrethor

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

// appendIfMissing ensures no duplicates in UsedBy slice
func appendIfMissing(slice []string, item string) []string {
	for _, existing := range slice {
		if existing == item {
			return slice
		}
	}
	return append(slice, item)
}

func printOrphaned(orphaned []OrphanedSecret, format string) error {
	switch format {
	case "json":
		b, _ := json.MarshalIndent(orphaned, "", " ")
		fmt.Println(string(b))
	case "yaml":
		b, _ := yaml.Marshal(orphaned)
		fmt.Println(string(b))
	}
	return nil
}
