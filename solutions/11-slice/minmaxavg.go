package main

import (
	"fmt"
)

// final return value is an indicator of success
func minmaxavg(nums []float64) (float64, float64, float64, bool) {
	if len(nums) == 0 { // handle empty slice or nil slice
		return 0.0, 0.0, 0.0, false
	}
	// "prime the pump"
	min, max, total := nums[0], nums[0], nums[0]

	// now look for nums larger or smaller and add to running total
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
		total += nums[i]
	}
	return min, max, total / float64(len(nums)), true
}

func main() {
	nums := []float64{} //54.2, 19.3, 23.4, -19.3, -54.2, -23.4}
	if min, max, avg, success := minmaxavg(nums); success {
		fmt.Println(min, max, avg)
	} else {
		fmt.Println("bad slice")
	}
}
