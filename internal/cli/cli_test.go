package cli

import "testing"

func TestCountry(t *testing.T) {
	Parse()

	if Country() != country {
		t.Fail()
	}
}

func TestFormat(t *testing.T) {
	Parse()

	if Format() != format {
		t.Fail()
	}
}
