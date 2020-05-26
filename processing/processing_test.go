package processing

import (
	"testing"

	"github.com/brunats/govid/providers"
)

func TestProcessing(t *testing.T) {
	var dataProviderFake providers.Data

	dataProviderFake.Deaths = 10
	dataProviderFake.Confirmed = 100

	data := Processing(dataProviderFake)

	if data.Processing.MortalityRate == 0 {
		t.Fail()
	}

	if data.Processing.MortalityRate != int((100*100)/10) {
		t.Fail()
	}
}
func TestCalculateMortalityRate(t *testing.T) {
	confirmed := 363211
	deaths := 22666
	mortalityRateExpected := int((confirmed * 100) / deaths)

	calculatedMortalityRate := calculateMortalityRate(confirmed, deaths)

	if calculatedMortalityRate != mortalityRateExpected {
		t.Fail()
	}
}
