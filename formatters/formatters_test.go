package formatters

import (
	"context"
	"testing"
)

type FormaterFake struct{}

func (f *FormaterFake) Receive(ctx context.Context)      {}
func (f *FormaterFake) Presentation(ctx context.Context) {}

func TestRegisterFormaters(t *testing.T) {
	Register(&FormaterFake{})

	if len(Formatters()) != 1 {
		t.Fail()
	}
}
