package nowsh

import (
	"errors"
	"io"
	"net/http"
)

type service interface {
	RequestCountries() (io.Reader, error)
}

type requestService struct{}

func (s *requestService) RequestCountries() (io.Reader, error) {
	resp, err := http.Get(NowShURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	return resp.Body, nil
}
