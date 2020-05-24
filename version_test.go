package main

import "testing"

func TestVersion(t *testing.T) {
	if Version != "1.0.0" {
		t.Fail()
	}
}
