package main

import (
	"testing"
)

func TestDistanceFromSun(t *testing.T) {
	venus := planetName{VENUS, 2000}

	expected := findDistance()

	actual := venus.distanceFromSun()

	if actual != expected {
		t.Error("Test has failed")
	}
}
