package formatters

import (
	"context"
	"testing"

	"github.com/brunats/govid/providers"
)

type FormaterFake struct{}

func (f *FormaterFake) Presentation(ctx context.Context, providersData []providers.Data) {}

func TestRegisterFormaters(t *testing.T) {
	Register(&FormaterFake{})

	if len(Formatters()) != 1 {
		t.Fail()
	}
}
