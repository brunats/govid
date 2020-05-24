package cli

import "flag"

var (
	country string = "ANY"
	format  string = "table"
)

// Parse cli flags
func Parse() {
	if !flag.Parsed() {
		flag.StringVar(&country, "country", country, "show only this country. Eg. [BR US ES ...]")
		flag.StringVar(&format, "format", format, "output format. Eg. [table json]")

		flag.Parse()
	}
}
