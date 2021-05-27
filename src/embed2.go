package main

import "fmt"

type Mountain struct {
	Name      string
	Elevation int
}

//START OMIT
type Climb struct {
	Mountain
	Climber string
}

func (m Mountain) HowBig() string { // HL
	if m.Elevation > 8800 {
		return "HUGE"
	}
	return "NORMAL"
}

func (c Climb) HowBig() string { // HL
	return "OVERRIDE!"
}

//END OMIT

func main() {
	c2 := Climb{
		Mountain{"K2", 8111},
		"Arjun Climber",
	}
	fmt.Println(c2.HowBig())
	fmt.Println(c2.Mountain.HowBig())
}
