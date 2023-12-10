package main

import (
	"strings"
)

type CalibrationValues struct {
	Values []string
}

// Calibrate will retrieve calibration values from each string and then sums it all.
// For each value, if only one integer is found, multiply it by 11 (ex: 7 => 77) and
// add it to sum.
// If two or more, it takes the first and last,
// do first * 10 + last (ex: 1 and 2 => 10 + 2 => 12) and adds it to sum.
func (c CalibrationValues) Calibrate() int {
	sum := 0
	repl := strings.NewReplacer(
		"one", "o1e", "two", "t2o", "three", "t3e", "four", "f4r", "five", "f5e", "six", "s6x", "seven", "s7n",
		"eight", "e8t", "nine", "n9e",
	)

	// Iterate over all values
	for _, str := range c.Values {
		rstr := repl.Replace(repl.Replace(str))
		first := int(rstr[strings.IndexAny(rstr, "123456789")] - '0')
		last := int(rstr[strings.LastIndexAny(rstr, "123456789")] - '0')
		sum = sum + (first*10 + last)
	}

	return sum
}
