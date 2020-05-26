package cli

import "strings"

// Country to filter
func Country() string {
	return strings.ToTitle(country)
}

// Format in output
func Format() string {
	return strings.ToTitle(format)
}
