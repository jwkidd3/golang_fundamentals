package main

import "fmt"

type Mountain struct {
	Name      string
	Elevation int
}

type Climb struct {
	Mountain
	Climber string
}

//START OMIT
func (m Mountain) HowBig() string { // HL
	if m.Elevation > 8800 {
		return "HUGE"
	}
	return "NORMAL"
}

// This method overrides the one above
func (c Climb) HowBig() string { // HL
	return "OVERRIDE!"
}

func main() {
	c2 := Climb{
		Mountain{"K2", 8111}, "Arjun Climber",
	}
	// will call overriding method
	fmt.Println(c2.HowBig()) // HL
	// will call overridden method
	fmt.Println(c2.Mountain.HowBig()) // HL
}

//END OMIT
