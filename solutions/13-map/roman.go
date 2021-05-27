package main

import (
	"fmt"
)

func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	// iterate thru the roman numeral
	// plug digits into map to get Arabic value
	// put them into an int slice of the appropriate size
	// [ 1000, 100, 50, 10 ]

	// create int slice, but make it one bigger than we need
	// ...we will see why below
	arabicVals := make([]int, len(numeral)+1)

	for index, digit := range numeral {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("bad digit: %c\n", digit)
			return 0
		}
	}
	fmt.Println("before:", arabicVals)

	// iterate through the int slice, looking for numbers which are less than
	// their neighbor–if you find one, make it negative, e.g., MCMXCIX = 1999
	// { 1000, 100, 1000, 10, 100, 1, 10 }
	// { 1000, -100, 1000, -10, 100, -1, 10 }
	// sum them up...1999

	total := 0
	// We no longer need to worry about falling off the end of the slice
	// since we have an extra value at the end which is 0. So we can just
	// iterate through the slice, and the last time we'll compare the final
	// Arabic value against the 0 that follows it, which is an extra comparison
	// but not a big deal.
	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}
	fmt.Println("after:", arabicVals)
	return total
}

func main() {
	fmt.Println(romanToArabic("MCLX"))
	fmt.Println(romanToArabic("MCMXCIX"))
	fmt.Println(romanToArabic("MCMZ"))
}
