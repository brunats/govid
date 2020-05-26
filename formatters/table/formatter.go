package table

import (
	"fmt"
	"strings"

	"github.com/brunats/govid/providers"
)

const (
	divider    = "+-------------------------------------------------------------------------------------------------------------------+\n"
	dividerAux = "---------------------------------------------------------------------------------------------------------------------\n"
)

// Formatter as table
type Formatter struct{}

// New table formatter
func New() *Formatter {
	return &Formatter{}
}

// Presentation of provider
func (f *Formatter) Presentation(providersData []*providers.Data) {
	var lines []string

	lineHeader := presentationHeader()
	lines = append(lines, divider, divider, lineHeader, divider)

	for _, provider := range providersData {
		lines = append(lines, presentationCountry(provider), dividerAux)
	}
	lines = append(lines, divider, divider)

	print(lines)
}

func presentationHeader() string {
	line := fmt.Sprintf(
		"+%32s %15s %15s %15s %15s %15s %1s+\n",
		"Country",
		"Confirmed",
		"Deaths",
		"Recovered",
		"Mortality Rate",
		"Source",
		" ",
	)

	return line
}

func presentationCountry(providerData *providers.Data) string {
	return fmt.Sprintf(
		"+%32s %15d %15d %15d %15d%% %15s %1s+\n",
		providerData.Country,
		providerData.Confirmed,
		providerData.Deaths,
		providerData.Recovered,
		providerData.Processing.MortalityRate,
		providerData.Provider,
		" ",
	)
}

func print(lines []string) {
	fmt.Println(strings.Join(lines, ""))
}
