package nowsh

import "testing"

func TestRequestServiceRequestCountries(t *testing.T) {
	service := &requestService{}

	_, err := service.RequestCountries()
	if err != nil {
		t.Fail()
	}
}
