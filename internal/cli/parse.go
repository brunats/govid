package cli

import "flag"

// Keys used in cli
const (
	CountryKey = "country"
	FormatKey  = "format"
)

var (
	country string = "ANY"
	format  string = "TABLE"
)

// Parse cli flags
func Parse() {
	if !flag.Parsed() {
		flag.StringVar(&country, CountryKey, country, "show only this country. Eg. [BR US ES ...]")
		flag.StringVar(&format, FormatKey, format, "output format. Eg. [table json]")

		flag.Parse()
	}
}
