package main

import (
	"reflect"
	"testing"
)

func TestStringFrequency(t *testing.T) {
	input := []string{"sachin", "rahul", "RAHUL", "rAhUl", "virat", "mahendra", "adam", "sachin", "rohit", "ajinkya", "sachin", "rahul", "virat", "sachin", "ricky"}

	expected := StringFrequency(input)

	actual := CallString()

	eq := reflect.DeepEqual(actual, expected)

	if !eq {
		t.Error("Test has failed!")

	}

}

/*import "reflect"
// m1 and m2 are the maps we want to compare
eq := reflect.DeepEqual(m1, m2)
if eq {
    fmt.Println("They're equal.")
} else {
    fmt.Println("They're unequal.")
}*/
