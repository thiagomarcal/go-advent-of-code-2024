package exercise2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1() {
	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer2.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	result := string(input)

	totalLines := len(strings.Split(result, "\n"))

	fmt.Printf("Total lines: %d\n", totalLines)

	arrResult := make([]string, totalLines)

	for _, line := range strings.Split(result, "\n") {
		if len(line) > 0 {

			var arr []int

			for _, res := range strings.Split(line, " ") {
				resToInt, _ := strconv.Atoi(strings.Trim(res, " "))
				arr = append(arr, resToInt)
			}

			counter := 0
			orientation := 0
			currentOrientation := 0

			c := checkErrors(arr, counter, orientation, currentOrientation)
			counter = c

			if counter > 0 {
				arrResult = append(arrResult, "Unsafe")
			} else {
				arrResult = append(arrResult, "Safe")
			}
		}

	}

	solution := 0
	for _, res := range arrResult {
		if res == "Safe" {
			solution++
		}
	}

	fmt.Printf("Exercise2 - Result: %v\n", solution)
	fmt.Printf("Took: %v\n", time.Since(timeInit))
}

func part2() {

	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer2-part2.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	result := string(input)

	totalLines := len(strings.Split(result, "\n"))

	arrResult := make([]string, totalLines)

	for _, line := range strings.Split(result, "\n") {
		if len(line) > 0 {

			var arr []int

			for _, res := range strings.Split(line, " ") {
				resToInt, _ := strconv.Atoi(strings.Trim(res, " "))
				arr = append(arr, resToInt)
			}

			counter := 0
			orientation := 0
			currentOrientation := 0

			c := checkErrors(arr, counter, orientation, currentOrientation)
			counter = c
			successCounter := 0

			// if there is an error, check if removing one element from the array solves the problem
			// otherwise, the array is safe
			if c > 0 {
				for i := range arr {

					// check every combination of the array without one element
					var newArr []int
					newCounter := 0
					newOrientation := 0
					newCurrentOrientation := 0

					for j, elem2 := range arr {
						if j != i {
							newArr = append(newArr, elem2)
						}
					}

					c2 := checkErrors(newArr, newCounter, newOrientation, newCurrentOrientation)
					if c2 == 0 {
						successCounter++
					}
				}

				// if at least one combination is safe, the array is safe
				if successCounter > 0 {
					arrResult = append(arrResult, "Safe")
				} else {
					arrResult = append(arrResult, "UnSafe")
				}
			} else {
				arrResult = append(arrResult, "Safe")
			}

		}
	}

	solution := 0
	for _, res := range arrResult {
		if res == "Safe" {
			solution++
		}
	}

	fmt.Printf("Exercise2 - Result part2: %v\n", solution)
	fmt.Printf("Took: %v\n", time.Since(timeInit))
}

func checkErrors(arr []int, counter int, orientation int, currentOrientation int) int {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] == 0 {
			counter++
			if counter == 1 {
				// if left and right are the same, the array is unsafe
				return counter
			}
		}

		if arr[i]-arr[i+1] < 0 {
			orientation = -1
		} else {
			orientation = 1
		}

		if currentOrientation == 0 {
			currentOrientation = orientation
		}

		if currentOrientation != orientation {
			counter++
			if counter == 1 {
				// if the orientation of the numbers is different, the array is unsafe
				return counter
			}
		}

		diff := math.Abs(float64(arr[i] - arr[i+1]))
		if diff > 3 {
			counter++
			if counter == 1 {
				// if the difference between the numbers is greater than 3, the array is unsafe
				return counter
			}
		}
	}
	return counter
}

func Solution() {
	part1()
	part2()
}
