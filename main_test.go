package main

/**
Execute w. "go test"

Before execution: 'go env -w GO111MODULE=off'
*/

import (
	"testing"
)

func TestSampleFunctionA(t *testing.T) {
	result := SampleFunctionA("winki")

	if result != "winki" {
		t.Error("Failure")
	}
}
