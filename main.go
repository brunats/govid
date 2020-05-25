package main

import (
	"context"
	"fmt"

	"github.com/brunats/govid/internal/cli"
)

func main() {
	cli.Parse()

	country := ctx().Value(cli.CountryKey).(string)
	format := ctx().Value(cli.FormatKey).(string)
	fmt.Println(format)
	fmt.Println(country)
}

func ctx() context.Context {
	return context.WithValue(
		context.WithValue(
			context.Background(),
			cli.CountryKey,
			cli.Country(),
		),
		cli.FormatKey,
		cli.Format(),
	)
}
