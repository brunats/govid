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
	dataFake := providers.Data{
		Provider:  "test",
		Error:     nil,
		Confirmed: 10000,
		Deaths:    1,
		Recovered: 9999,
		Country:   "Brazil",
	}

	line := presentationCountry(dataFake)

	lineExpected := fmt.Sprintf("+%35s %15d %15d %15d %1s+\n", dataFake.Country, dataFake.Confirmed, dataFake.Deaths, dataFake.Recovered, " ")
	if line != lineExpected {
		t.Fail()
	}
}

func TestPresentationHeader(t *testing.T) {
	providerName := "dataFake"
	var linesExpected []string
	linesExpected = append(linesExpected, fmt.Sprintf("+ Source: %s %70s+\n", providerName, " "))
	linesExpected = append(linesExpected, fmt.Sprintf("+%35s %15s %15s %15s %1s+\n", "Country", "Confirmed", "Deaths", "Recovered", " "))

	headerLines := presentationHeader(providerName)

	for index, lineExpected := range linesExpected {
		if headerLines[index] != lineExpected {
			t.Fail()
		}
	}
}
