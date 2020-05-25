package formatters

import (
	"context"

	"github.com/brunats/govid/providers"
)

var formatters []Formatter

// Formatter interface
type Formatter interface {
	Presentation(ctx context.Context, providersData []providers.Data)
}

// Register formatter
func Register(formatter Formatter) {
	formatters = append(formatters, formatter)
}

// Formatters returns registered formatters
func Formatters() []Formatter {
	return formatters
}
