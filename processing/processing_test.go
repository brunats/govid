package processing

import (
	"testing"

	"github.com/brunats/govid/providers"
)

func TestProcessing(t *testing.T) {
	dataProviderFake := &providers.Data{Deaths: 10, Confirmed: 100}

	Processing(dataProviderFake)

	if dataProviderFake.Processing.MortalityRate == 0 {
		t.Fail()
	}

	if dataProviderFake.Processing.MortalityRate != int((10*100)/100) {
		t.Fail()
	}
}
func TestCalculateMortalityRateWithDeaths(t *testing.T) {
	confirmed := 363211
	deaths := 22666
	mortalityRateExpected := int((deaths * 100) / confirmed)

	calculatedMortalityRate := calculateMortalityRate(confirmed, deaths)

	if calculatedMortalityRate != mortalityRateExpected {
		t.Fail()
	}
}

func TestCalculateMortalityRateWithuotDeaths(t *testing.T) {
	confirmed := 363211
	deaths := 0
	mortalityRateExpected := 0

	calculatedMortalityRate := calculateMortalityRate(confirmed, deaths)

	if calculatedMortalityRate != mortalityRateExpected {
		t.Fail()
	}
}
