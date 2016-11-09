/*

@author rakesh.narang@nagarro.com

algorithm

1. declare array of size 9 (0-9) to store the recurrence of digits
2. hard coded input
3. store in relevant positions
4. display the frequency

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	input := [15]int{1, 3, 4, 5, 2, 3, 4, 5, 6, 7, 8, 3, 4, 5, 6}

	countNumbers(input)

}

func countNumbers(input [15]int) {

	var store [10]int

	for i := 0; i < len(input); i++ {

		store[input[i]]++

	}

	fmt.Println(PrintFrequency(store))

}

func PrintFrequency(store [10]int) (string, error) {
	var retVal string

	for i := 0; i < len(store); i++ {
		retVal += strconv.Itoa(store[i])
	}

	return retVal

}
