package formatters

import "context"

var formatters []Formatter

// Formatter interface
type Formatter interface {
	Receive(ctx context.Context)
	Presentation(ctx context.Context)
}

// Register formatter
func Register(formatter Formatter) {
	formatters = append(formatters, formatter)
}

// Formatters returns registered formatters
func Formatters() []Formatter {
	return formatters
}
