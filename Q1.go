package main

import (
	"fmt"
)

type Planet int

func (p Planet) String() string {
	switch p {
	case MERCURY:
		return "Mercury"
	case VENUS:
		return "Venus"
	case EARTH:
		return "Earth"
	case MARS:
		return "Mars"
	case JUPITER:
		return "Jupiter"
	case SATURN:
		return "Saturn"
	case URANUS:
		return "Uranus"
	case NEPTUNE:
		return "Neptune"
	case PLUTO:
		return "Pluto"
	}
	return ""
}

const (
	MERCURY Planet = iota
	VENUS
	EARTH
	MARS
	JUPITER
	SATURN
	URANUS
	NEPTUNE
	PLUTO
)

func main() {
	fmt.Printf("%d\t%s\n", URANUS, URANUS)
}
