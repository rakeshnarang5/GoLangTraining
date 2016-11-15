package main

import (
	"fmt"
)

type Planet interface {
	Name() string
	Mass() int64
}

type myPlanet int

const (
	MERCURY myPlanet = iota + 1
	VENUS
	EARTH
	MARS
	JUPITER
	SATURN
	URANUS
	NEPTUNE
	PLUTO
)

func (p myPlanet) String() string {
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

type planetName struct {
	obj myPlanet
	pos int
}

func (p *planetName) Name() string {
	retVal := fmt.Sprintf("%s", p.obj)
	return retVal
}

func (p *planetName) Mass() int64 {
	switch p.obj {
	case MERCURY:
		return 10
	case VENUS:
		return 100
	case EARTH:
		return 1000
	case MARS:
		return 10000
	case JUPITER:
		return 100000
	case SATURN:
		return 1000000
	case URANUS:
		return 10000000
	case NEPTUNE:
		return 100000000
	case PLUTO:
		return 1000000000
	}
	return 0
}

func (p *planetName) distanceFromSun() int {
	return 10000 - p.pos
}

func findDistance() int {
	venus := planetName{VENUS, 2000}
	return venus.distanceFromSun()
}

func dummy(p planetName) {
	fmt.Println(p.Name())
	fmt.Println(p.Mass())
	fmt.Println(p.distanceFromSun())
}

func main() {
	venus := planetName{VENUS, 2000}
	dummy(venus)
}
