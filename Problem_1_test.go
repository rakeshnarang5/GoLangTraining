package main

import "testing"

func TestPrintFrequency(t *testing.T) {
	expected := "0113332110"
	store := [10]int{0, 1, 1, 3, 3, 3, 2, 1, 1, 1}
	actual := PrintFrequency(store)

	if actual != expected {
		t.Error("Test failed!")
	}
}
