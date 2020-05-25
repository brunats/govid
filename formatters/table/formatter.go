package table

import (
	"context"
	"fmt"

	"github.com/brunats/govid/formatters"
	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/providers"
)

const divider = "+-------------------------------------------------------------------------------------+\n"
const dividerAux = "---------------------------------------------------------------------------------------\n"

type formatter struct{}

// New table formatter
func New() formatters.Formatter {
	return &formatter{}
}

// Presentation of provider
func (f *formatter) Presentation(ctx context.Context, providersData []providers.Data) {
	if ctx.Value(cli.FormatKey).(string) != "table" {
		return
	}

	var lines []string

	linesOfHeader := presentationHeader(providersData[0].Provider)
	lines = append(lines, divider, divider)
	lines = append(lines, linesOfHeader...)
	lines = append(lines, divider)

	for _, provider := range providersData {
		lines = append(lines, presentationCountry(provider), dividerAux)
	}
	lines = append(lines, divider, divider)

	print(lines)
}

func presentationHeader(provider string) []string {
	var lines []string
	lines = append(lines, fmt.Sprintf("+ Source: %s %70s+\n", provider, " "))
	lines = append(lines, fmt.Sprintf("+%35s %15s %15s %15s %1s+\n", "Country", "Confirmed", "Deaths", "Recovered", " "))

	return lines
}

func presentationCountry(providerData providers.Data) string {
	return fmt.Sprintf("+%35s %15d %15d %15d %1s+\n", providerData.Country, providerData.Confirmed, providerData.Deaths, providerData.Recovered, " ")
}

func print(lines []string) {

}
