package nowsh

import (
	"bufio"
	"io"
	"os"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	provider := New()

	if provider == nil {
		t.Fail()
	}
}

type serviceFake interface {
	RequestCountries() (io.Reader, error)
}

type requestServiceFake struct{}

func (s *requestServiceFake) RequestCountries() (io.Reader, error) {
	file, _ := os.Open("allCountries.json")
	reader := bufio.NewReader(file)

	return reader, nil
}

func TestRequestCountries(t *testing.T) {
	wg := sync.WaitGroup{}

	provider := &provider{}
	provider.wg = wg
	provider.service = &requestServiceFake{}

	provider.requestCountries()

	if len(provider.response) != 5 {
		t.Fail()
	}
}
