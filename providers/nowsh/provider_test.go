package nowsh

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

type serviceFake interface {
	RequestCountries() (io.Reader, error)
}

type requestServiceFakeWithoutErrors struct{}
type requestServiceFakeWithErrors struct{}

func (s *requestServiceFakeWithoutErrors) RequestCountries() (io.Reader, error) {
	file, _ := os.Open("allCountries.json")
	reader := bufio.NewReader(file)

	return reader, nil
}

func (s *requestServiceFakeWithErrors) RequestCountries() (io.Reader, error) {
	err := fmt.Errorf("fake error")

	return nil, err
}

func TestNew(t *testing.T) {
	provider := New()

	if provider == nil {
		t.Fail()
	}
}

func TestRequestCountries(t *testing.T) {
	wg := sync.WaitGroup{}

	provider := &provider{wg: wg, service: &requestServiceFakeWithoutErrors{}}
	provider.requestCountries()

	if len(provider.response) != 5 {
		t.Fail()
	}
}

func TestRequestCountriesWithError(t *testing.T) {
	wg := sync.WaitGroup{}

	provider := &provider{wg: wg, service: &requestServiceFakeWithErrors{}}

	provider.requestCountries()

	if len(provider.response) > 1 {
		t.Fail()
	}

	if provider.response[0].Error == nil {
		t.Fail()
	}
}
