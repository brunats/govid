package nowsh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/brunats/govid/internal/cli"
	"github.com/brunats/govid/providers"
)

type provider struct {
	wg       sync.WaitGroup
	response []providers.Data
}

type countryInfo struct {
	Country   string    `json:"country"`
	Cases     int       `json:"cases"`
	Confirmed int       `json:"confirmed"`
	Recovered int       `json:"recovered"`
	Deaths    int       `json:"deaths"`
	UpdatedAt time.Time `json:"updated_at"`
}

type responseForm struct {
	CountryInfos []countryInfo `json:"data"`
}

// New nowsh provider
func New() providers.Provider {
	wg := sync.WaitGroup{}

	return &provider{wg: wg}
}

func (p *provider) Request(ctx context.Context) {
	p.wg.Add(1)
	defer p.wg.Done()

	if ctx.Value(cli.CountryKey).(string) == "ANY" {
		requestCountries(ctx)
	}
}

func (p *provider) Wait() {
	p.wg.Wait()
}

func (p *provider) Response() []providers.Data {
	return p.response
}

func (p *provider) requestCountries(context.Context) {
	resp, err := http.Get("https://covid19-brazil-api.now.sh/api/report/v1/countries")
	if err != nil {
		p.appendErrorResponse(err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		form := &responseForm{}
		err := json.NewDecoder(resp.Body).Decode(form)
		if err != nil {
			p.appendErrorResponse(err)
		}

		for _, info := range form.CountryInfos {
			p.appendResponse(info)
		}
	} else {
		p.appendErrorResponse(errors.New("bad request"))
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
