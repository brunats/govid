package main

import (
	"context"

	"github.com/brunats/govid/formatters"
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

	var dataProviders []providers.Data
	for _, provider := range providers.Providers() {
		dataProviders = append(dataProviders, provider.Response()...)
	}

	formatter := formatters.Selection(ctx)
	formatter.Presentation(dataProviders)
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
