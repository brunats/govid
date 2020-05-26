package formatters

import (
	"context"
	"reflect"
	"testing"

	"github.com/brunats/govid/formatters/table"
	"github.com/brunats/govid/providers"
)

type FormaterFake struct{}

var ctxFake = context.WithValue(
	context.WithValue(
		context.Background(),
		"country",
		"ANY",
	),
	"format",
	"table",
)

func (f *FormaterFake) Presentation(providersData []providers.Data) {}

func TestRegisterFormatters(t *testing.T) {
	Register(&FormaterFake{})

	if len(Formatters()) != 1 {
		t.Fail()
	}
}

func TestSelectFormatter(t *testing.T) {
	Register(&FormaterFake{})

	formatter := Selection(ctxFake)

	if formatter == nil {
		t.Fail()
	}

	typeOfObjectTable := reflect.TypeOf(table.New())
	if reflect.TypeOf(formatter) != typeOfObjectTable {
		t.Fail()
	}

}
