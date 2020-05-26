package json

import (
	"encoding/json"
	"fmt"

	"github.com/brunats/govid/providers"
)

// Formatter as table
type Formatter struct{}

// New json formatter
func New() *Formatter {
	return &Formatter{}
}

// Presentation of provider
func (f *Formatter) Presentation(providersData []*providers.Data) {
	for _, provider := range providersData {
		fmt.Println(convertProviderDataToString(provider))
	}
}

func convertProviderDataToString(providersData *providers.Data) string {
	message, err := json.Marshal(providersData)

	if err != nil {
		return fmt.Sprintf("{ Data: { Error: %s } }", err)
	}

	return string(message)
}
