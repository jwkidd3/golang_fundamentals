package main

import "fmt"

// START OMIT
import "math"

const maxnum = 100

func main() {
	numbers := make([]int, maxnum+1)
	numbers[0], numbers[1] = 0, 0

	for num := 2; num <= maxnum; num++ {
		numbers[num] = num
	}
	for index := 2; index <= int(math.Sqrt(float64(maxnum))+1); index++ {
		if numbers[index] == 0 {
			continue
		}
		for multiple := 2 * numbers[index]; multiple <= maxnum; multiple += index {
			numbers[multiple] = 0
		}
	}
	for _, value := range numbers {
		if value > 0 {
			fmt.Println(value)
		}
	}
}

// END OMIT
