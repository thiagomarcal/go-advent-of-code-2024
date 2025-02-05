package exercise1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	input1   int
	input2   int
	distance int
}

func part1() {
	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer1.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	result := string(input)

	totalLines := len(strings.Split(result, "\n"))

	fmt.Printf("Total lines: %d\n", totalLines)

	arr1 := make([]int, totalLines)
	arr2 := make([]int, totalLines)
	final := make([]Input, totalLines)

	totalDistance := 0

	for i, line := range strings.Split(result, "\n") {
		splitted := strings.Split(line, "   ")

		if len(splitted) > 1 {

			input1 := splitted[0]

			res1, err := strconv.Atoi(strings.Trim(input1, " "))
			if err != nil {
				fmt.Printf("Error converting input1 to Integer: %v\n", err)
			}

			input2 := splitted[1]
			res2, err := strconv.Atoi(strings.Trim(input2, " "))
			if err != nil {
				fmt.Printf("Error converting input2 to Integer: %v\n", err)
			}

			arr1[i] = res1
			arr2[i] = res2

		}
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < totalLines; i++ {
		distance := 0
		if arr1[i] > arr2[i] {
			distance = arr1[i] - arr2[i]
		} else {
			distance = arr2[i] - arr1[i]
		}

		final[i] = Input{input1: arr1[i], input2: arr2[i], distance: distance}
		totalDistance += final[i].distance
	}

	took := time.Since(timeInit)

	fmt.Printf("Exercise1 - Result: %v\n", totalDistance)

	fmt.Printf("Took: %v\n", took)

}

func part2() {
	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer1-part2.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	result := string(input)

	totalLines := len(strings.Split(result, "\n"))

	fmt.Printf("Total lines: %d\n", totalLines)

	arrLeft := make([]int, totalLines)
	hashmapRight := make(map[int]int)

	for i, line := range strings.Split(result, "\n") {
		if len(line) > 0 {

			var arr []int

			for _, res := range strings.Split(line, "   ") {
				resToInt, _ := strconv.Atoi(strings.Trim(res, " "))
				arr = append(arr, resToInt)
			}

			arrLeft[i] = arr[0]
			numFreq := arr[1]

			if _, ok := hashmapRight[numFreq]; ok {
				hashmapRight[numFreq]++
			} else {
				hashmapRight[numFreq] = 1
			}
		}
	}

	solution := 0

	for _, num := range arrLeft {
		if _, ok := hashmapRight[num]; ok {
			solution += hashmapRight[num] * num
		}
	}

	fmt.Printf("Exercise1 - Result Part2: %v\n", solution)

	fmt.Printf("Took: %v\n", time.Since(timeInit))
}

func Solution() {
	part1()
	part2()
}
