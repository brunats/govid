package nowsh

import (
	"errors"
	"io"
	"net/http"
)

type service interface {
	RequestCountries() (io.Reader, error)
	RequestCountry(countryName string) (io.Reader, error)
}

type requestService struct{}

func (s *requestService) RequestCountries() (io.Reader, error) {
	resp, err := http.Get(NowShURLAll)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	return resp.Body, nil
}

func (s *requestService) RequestCountry(countryName string) (io.Reader, error) {
	resp, err := http.Get(NowShURLOn + countryName)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	return resp.Body, nil
}
