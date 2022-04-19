package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var decMeasurement int
var incMeasurement int
var lastMeasurement int = -1
var measurement int
var measurements = []int{-1, -1, -1}
var path string = "input.txt"
var slidingWindowIndex int
var slidingWindowMod int

func main() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open file %s, %v", path, err))
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		measurement, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("(invalid data")
			continue
		}

		measurements[2] = measurements[1]
		measurements[1] = measurements[0]
		measurements[0] = measurement

		if measurements[2] != -1 {
			currentMeasurement := sumOverArray(measurements)

			if lastMeasurement == -1 {
				fmt.Println("(N/A - no previous sum)")
			} else if currentMeasurement == lastMeasurement {
				fmt.Println("(no change)")
			} else if currentMeasurement > lastMeasurement {
				fmt.Println("(increased)")
				incMeasurement++
			} else if currentMeasurement < lastMeasurement {
				fmt.Println("(decreased)")
				decMeasurement++
			}

			lastMeasurement = currentMeasurement
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Increased three-measurements: %v\n", incMeasurement)
	fmt.Printf("Decreased three-measurements: %v\n", decMeasurement)
}

func sumOverArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
