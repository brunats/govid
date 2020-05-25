package main

import (
	"context"
	"fmt"

	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/providers"
	"github.com/brunats/govid/providers/nowsh"
)

func main() {
	cli.Parse()

	// Register providers
	providers.Register(nowsh.New())

	// Request providers
	ctx := ctx()

	for _, provider := range providers.Providers() {
		provider.Request(ctx)
	}

	for _, provider := range providers.Providers() {
		provider.Wait()
	}

	for _, provider := range providers.Providers() {
		fmt.Println(provider.Response())
	}
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
