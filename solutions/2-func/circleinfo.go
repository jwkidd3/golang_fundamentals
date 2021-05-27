// Write a function called `circleinfo` which accepts a float64 radius
// and returns two values, the area of the circle, and the circumference
package main

import (
	"fmt"
	"math"
)

func circleinfo(r float64) (float64, float64) {
	return math.Pi * r * r, math.Pi * r * 2
}

func main() {
	area, circ := circleinfo(3.5)
	fmt.Printf("area = %.2f, circumference = %.2f\n", area, circ)
}
