package nowsh

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/providers"
)

// NowShURL used to call api
const NowShURL = "https://covid19-brazil-api.now.sh/api/report/v1/countries"

type provider struct {
	wg       sync.WaitGroup
	response []providers.Data
	service  service
}

// New nowsh provider
func New() providers.Provider {
	wg := sync.WaitGroup{}

	return &provider{wg: wg, service: &requestService{}}
}

func (p *provider) Request(ctx context.Context) {
	p.wg.Add(1)
	defer p.wg.Done()

	if ctx.Value(cli.CountryKey).(string) == "ANY" {
		p.requestCountries()
	}
}

func (p *provider) Wait() {
	p.wg.Wait()
}

func (p *provider) Response() []providers.Data {
	return p.response
}

func (p *provider) requestCountries() {
	reader, err := p.service.RequestCountries()

	if err == nil {
		form := &responseForm{}
		err := json.NewDecoder(reader).Decode(form)
		if err != nil {
			p.appendErrorResponse(err)
		}

		for _, info := range form.CountryInfos {
			p.appendResponse(info)
		}
	} else {
		p.appendErrorResponse(err)
	}
}

func (p *provider) appendResponse(info countryInfo) {
	p.response = append(
		p.response,
		providers.Data{
			Provider:  "nowsh",
			Error:     nil,
			Confirmed: info.Confirmed,
			Deaths:    info.Deaths,
			Recovered: info.Recovered,
			Country:   info.Country,
		},
	)
}

func (p *provider) appendErrorResponse(err error) {
	p.response = append(
		p.response,
		providers.Data{
			Provider: "nowsh",
			Error:    fmt.Errorf("nowsh provider: %w", err),
		},
	)
}
