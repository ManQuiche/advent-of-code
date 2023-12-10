package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	filePathPtr := flag.String("f", "./puzzle.txt", "path to calibration values file")
	values, err := LoadFromFile(*filePathPtr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(values.Calibrate())
}
