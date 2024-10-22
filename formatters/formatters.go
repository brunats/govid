package formatters

import (
	"context"

	"github.com/brunats/govid/formatters/json"
	"github.com/brunats/govid/formatters/table"
	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/providers"
)

var formatters []Formatter

// Formatter interface
type Formatter interface {
	Presentation(providersData []*providers.Data)
}

// Register formatter
func Register(formatter Formatter) {
	formatters = append(formatters, formatter)
}

// Selection a provider
func Selection(ctx context.Context) Formatter {
	formatTable := "TABLE"
	formatJSON := "JSON"

	switch ctx.Value(cli.FormatKey).(string) {
	case formatTable:
		return table.New()

	case formatJSON:
		return json.New()

	default:
		return table.New()
	}
}

// Formatters returns registered formatters
func Formatters() []Formatter {
	return formatters
}
