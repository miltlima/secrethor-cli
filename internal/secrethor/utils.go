package secrethor

// appendIfMissing ensures no duplicates in UsedBy slice
func appendIfMissing(slice []string, item string) []string {
	for _, existing := range slice {
		if existing == item {
			return slice
		}
	}
	return append(slice, item)
}
