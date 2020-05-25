package nowsh

import "testing"

func TestNew(t *testing.T) {
	provider := New()

	if provider == nil {
		t.Fail()
	}
}

type serviceFake interface {
	RequestCountriesFake() (io.Reader, error)
}

func RequestCountriesFake() (io.Reader, error) {
	resp, err := http.Get(NowShURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	return resp.Body, nil
}


func TestRequest(t *testing.T) {
	
}