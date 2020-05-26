package json

import (
	"testing"

	"github.com/brunats/govid/providers"
)

func TestNew(t *testing.T) {
	formatter := New()

	if formatter == nil {
		t.Fail()
	}
}

func TestConvertProvidersDataToString(t *testing.T) {
	var mortalityRateFake providers.Processing
	mortalityRateFake.MortalityRate = 2
	dataFake := &providers.Data{
		Provider:   "test",
		Error:      nil,
		Confirmed:  10000,
		Deaths:     1,
		Recovered:  9999,
		Country:    "Brazil",
		Processing: mortalityRateFake,
	}

	stringExpected := "{\"Confirmed\":10000,\"Deaths\":1,\"Recovered\":9999,\"Error\":null,\"Provider\":\"test\",\"Country\":\"Brazil\",\"Processing\":{\"MortalityRate\":2}}"

	stringOfProvidersData := convertProviderDataToString(dataFake)

	if stringOfProvidersData != stringExpected {
		t.Fail()
	}
}
