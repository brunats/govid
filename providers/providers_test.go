package providers

import (
	"context"
	"testing"
)

type ProviderFake struct {
	ResponseMock []*Data
}

func (p *ProviderFake) Request(ctx context.Context) {}
func (p *ProviderFake) Wait()                       {}

func (p *ProviderFake) Response() []*Data {
	return p.ResponseMock
}

func TestRegisterProviders(t *testing.T) {
	Register(&ProviderFake{})

	if len(Providers()) != 1 {
		t.Fail()
	}
}
