package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var path string = "input.txt"

func main() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open file %s, %v", path, err))
		os.Exit(1)
	}
	defer file.Close()

	var lastMeasurement int = -1
	var incMeasurement int = 0
	var decMeasurement int = 0
	var measurement int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		measurement, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("(invalid data")
			continue
		}

		if lastMeasurement == -1 {
			fmt.Println("(N/A - no previous measurement)")
		} else if measurement == lastMeasurement {
			fmt.Println("(no difference)")
		} else if measurement > lastMeasurement {
			fmt.Println("(increased)")
			incMeasurement++
		} else if measurement < lastMeasurement {
			fmt.Println("(decreased)")
			decMeasurement++
		}

		lastMeasurement = measurement
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Increased measurements: %v\n", incMeasurement)
	fmt.Printf("Decreased measurements: %v\n", decMeasurement)
}
