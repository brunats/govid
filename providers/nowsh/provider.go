package nowsh

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/providers"
)

// NowShURLAll used to call api for all countries
const NowShURLAll = "https://covid19-brazil-api.now.sh/api/report/v1/countries"

// NowShURLOn used to call api for all countries
const NowShURLOn = "https://covid19-brazil-api.now.sh/api/report/v1/"

type provider struct {
	wg       sync.WaitGroup
	response []*providers.Data
	service  service
}

// New nowsh provider
func New() providers.Provider {
	wg := sync.WaitGroup{}
	wg.Add(1)

	return &provider{wg: wg, service: &requestService{}}
}

func (p *provider) Request(ctx context.Context) {
	defer p.wg.Done()
	country := ctx.Value(cli.CountryKey).(string)

	if country == "ANY" {
		p.requestCountries()
	} else {
		p.requestCountry(country)
	}
}

func (p *provider) Wait() {
	p.wg.Wait()
}

func (p *provider) Response() []*providers.Data {
	return p.response
}

func (p *provider) requestCountries() {
	reader, err := p.service.RequestCountries()

	if err == nil {
		form := &responseFormAll{}
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

func (p *provider) requestCountry(countryAbbreviation string) {
	reader, err := p.service.RequestCountry(getCountryFullName(countryAbbreviation))

	if err == nil {
		form := &responseFormOne{}
		err := json.NewDecoder(reader).Decode(form)
		if err != nil {
			p.appendErrorResponse(err)
		}

		p.appendResponse(form.CountryInfo)
	} else {
		p.appendErrorResponse(err)
	}
}

func getCountryFullName(abbreviation string) string {
	return CountryCode[strings.ToTitle(abbreviation)]
}

func (p *provider) appendResponse(info countryInfo) {
	p.response = append(
		p.response,
		&providers.Data{
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
		&providers.Data{
			Provider: "nowsh",
			Error:    fmt.Errorf("nowsh provider: %w", err),
		},
	)
}
