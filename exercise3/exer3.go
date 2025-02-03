package exercise3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func extractMulQuery(str string) []string {
	pattern := "mul\\([0-9]+,[0-9]+\\)"
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}

func multiply(str string) int {
	pattern := "[0-9]+,[0-9]+"
	re := regexp.MustCompile(pattern)
	results := re.FindAllString(str, -1)

	num1 := 0
	num2 := 0

	for _, res := range results {
		nums := strings.Split(res, ",")
		num1, _ = strconv.Atoi(nums[0])
		num2, _ = strconv.Atoi(nums[1])
	}

	return num1 * num2
}

func removeDont(str string) string {
	pattern := "don't\\((.*?)do\\("
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(str, "")
}


func part1() {
	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer3.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	result := strings.TrimSpace(string(input))
	sum := 0

	results := extractMulQuery(result)

	for _, query := range results {
		sum += multiply(query)
	}

	fmt.Printf("Exercise3 - Result: %v\n", sum)
	fmt.Printf("Took: %v\n", time.Since(timeInit))
}


func part2() {
	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer3-part2.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	trimmedInput := strings.TrimSpace(string(input))
	resultTrimmedAndWithoutNewLines := strings.ReplaceAll(trimmedInput, "\n", "")

	sum := 0

	revisedResult := removeDont(resultTrimmedAndWithoutNewLines)

	results := extractMulQuery(revisedResult)

	for _, query := range results {
		sum += multiply(query)
	}

	fmt.Printf("Exercise3 - Result Part2: %v\n", sum)
	fmt.Printf("Took: %v\n", time.Since(timeInit))

}

func Solution() {
	part1()
	part2()
}
