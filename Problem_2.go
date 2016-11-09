/*

@author rakesh.narang@nagarro.com

Algorithm:-

1. hard code an slices of string type of 15 values
2. create a map of key:string val:int
3. traverse the hard coded slices and perform following
	check if key exists
	if exists,
	get value, increment it, set value
	else,
	add key
4. output map

*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	m := CallString()
	fmt.Println(m)
}

func CallString() map[string]int {
	input := []string{"sachin", "rahul", "RAHUL", "rAhUl", "virat", "mahendra", "adam", "sachin", "rohit", "ajinkya", "sachin", "rahul", "virat", "sachin", "ricky"}

	m := StringFrequency(input)

	return m
}

func StringFrequency(input []string) map[string]int {
	m := make(map[string]int)

	for i := 0; i < len(input); i++ {

		val := strings.ToLower(input[i])

		_, prs := m[val]

		if prs == false {

			m[val] = 1

		} else {

			v1 := m[val]
			v1++
			m[val] = v1

		}
	}

	return m
}

/*func RetValString( m map[string]int) string {

	var retVal string

	for _, v := range m {
		retVal += strconv.Itoa(v)
	}

	return retVal
}*/
