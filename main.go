package main

import (
	"context"

	"github.com/brunats/govid/formatters"
	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/processing"
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

	// Awaiting completion
	for _, provider := range providers.Providers() {
		provider.Wait()
	}

	// Merge results
	var dataProviders []*providers.Data
	for _, provider := range providers.Providers() {
		dataProviders = append(dataProviders, provider.Response()...)
	}

	// Processing
	for _, data := range dataProviders {
		processing.Processing(data)
	}

	// Results
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
