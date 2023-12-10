package main

import (
	"regexp"
	"strconv"
)

type CalibrationValues struct {
	Values []string
}

// Calibrate will retrieve calibration values from each string and then sums it all.
// For each value, if only one integer is found, multiply it by 11 (ex: 7 => 77) and
// add it to sum. If two or more, it will take the first and last,
// do first * 10 + last (ex: 1 and 2 => 10 + 2 => 12) and add it to sum.c
func (c CalibrationValues) Calibrate() int {
	r, _ := regexp.Compile("([1-9])")
	sum := 0

	// Iterate over all values
	for _, str := range c.Values {
		numbers := r.FindAllString(str, -1)
		switch len(numbers) {
		case 0:
			continue
		case 1:
			value, _ := strconv.Atoi(numbers[0])
			sum = sum + value*11
		default:
			first, _ := strconv.Atoi(numbers[0])
			last, _ := strconv.Atoi(numbers[len(numbers)-1])
			sum = sum + (first*10 + last)
		}
	}

	return sum
}
