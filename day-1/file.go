package main

import (
	"fmt"
	"os"
	"strings"
)

func loadFromFile(file string) (CalibrationValues, error) {
	values, err := os.ReadFile(file)
	if err != nil {
		return CalibrationValues{}, fmt.Errorf("could not load values: %w", err)
	}
	valuesStr := strings.Split(string(values), "\n")

	return CalibrationValues{Values: valuesStr}, nil
}
