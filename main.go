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

	ctx := ctx()


	// Request providers
	for _, provider := range providers.Providers() {
		go provider.Request(ctx)
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
