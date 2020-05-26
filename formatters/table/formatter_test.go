package table

import (
	"fmt"
	"testing"

	"github.com/brunats/govid/providers"
)

func TestNew(t *testing.T) {
	formatter := New()

	if formatter == nil {
		t.Fail()
	}
}

func TestPresentationCountry(t *testing.T) {
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

	line := presentationCountry(dataFake)

	lineExpected := fmt.Sprintf(
		"+%32s %15d %15d %15d %15d%% %15s %1s+\n",
		dataFake.Country,
		dataFake.Confirmed,
		dataFake.Deaths,
		dataFake.Recovered,
		dataFake.Processing.MortalityRate,
		dataFake.Provider,
		" ",
	)
	if line != lineExpected {
		t.Fail()
	}
}

func TestPresentationHeader(t *testing.T) {
	lineExpected := fmt.Sprintf(
		"+%32s %15s %15s %15s %15s %15s %1s+\n",
		"Country",
		"Confirmed",
		"Deaths",
		"Recovered",
		"Mortality Rate",
		"Source",
		" ",
	)

	headerLines := presentationHeader()

	if headerLines != lineExpected {
		t.Fail()
	}
}
