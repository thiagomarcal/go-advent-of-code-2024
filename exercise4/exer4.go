package exercise4

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func countHorizontal(grid *[][]string, searchTerm string) int {
	pattern := searchTerm

	inversePattern := ""
	for i := len(searchTerm) - 1; i >= 0; i-- {
		inversePattern += string(searchTerm[i])
	}

	i := 0
	count := 0

	for i < len(*grid) {
		line := ""
		for j := 0; j < len((*grid)[i]); j++ {
			line += (*grid)[i][j]
		}
		// find word in line with regex
		re := regexp.MustCompile(pattern)
		results := re.FindAllString(line, -1)
		count += len(results)

		ire := regexp.MustCompile(inversePattern)
		iresults := ire.FindAllString(line, -1)
		count += len(iresults)

		i++
	}

	return count
}

func countVertical(grid *[][]string, searchTerm string) int {

	pattern := searchTerm
	inversePattern := ""

	for i := len(searchTerm) - 1; i >= 0; i-- {
		inversePattern += string(searchTerm[i])
	}

	i := 0
	count := 0

	for i < len(*grid) {
		line := ""
		for j := 0; j < len((*grid)[i]); j++ {
			line += (*grid)[j][i]
		}
		// find word in line with regex
		re := regexp.MustCompile(pattern)
		results := re.FindAllString(line, -1)
		count += len(results)

		ire := regexp.MustCompile(inversePattern)
		iresults := ire.FindAllString(line, -1)
		count += len(iresults)

		i++
	}

	return count
}

func countDiagonal(grid *[][]string, searchTerm string) int {

	pattern := searchTerm

	// inverse searchTerm
	inversePattern := ""
	for i := len(searchTerm) - 1; i >= 0; i-- {
		inversePattern += string(searchTerm[i])
	}

	count := 0

	rowLength := len((*grid)[0])

	re := regexp.MustCompile(pattern)
	ire := regexp.MustCompile(inversePattern)

	var mainDiags []string
	for col := 0; col < rowLength; col++ {
		i, j := 0, col
		diag := ""
		for i < rowLength && j < rowLength {
			diag += (*grid)[i][j]
			i++
			j++
		}
		mainDiags = append(mainDiags, diag)
	}

	for row := 1; row < rowLength; row++ {
		i, j := row, 0
		diag := ""
		for i < rowLength && j < rowLength {
			diag += (*grid)[i][j]
			i++
			j++
		}
		mainDiags = append(mainDiags, diag)
	}

	for _, diag := range mainDiags {

		results := re.FindAllString(diag, -1)
		count += len(results)

		iresults := ire.FindAllString(diag, -1)
		count += len(iresults)

	}

	var antiDiags []string

	for col := 0; col < rowLength; col++ {
		i, j := 0, col
		diag := ""
		for i < rowLength && j >= 0 {
			diag += (*grid)[i][j]
			i++
			j--
		}
		antiDiags = append(antiDiags, diag)
	}

	for row := 1; row < rowLength; row++ {
		i, j := row, rowLength-1
		diag := ""
		for i < rowLength && j >= 0 {
			diag += (*grid)[i][j]
			i++
			j--
		}
		antiDiags = append(antiDiags, diag)
	}

	for _, diag := range antiDiags {

		results := re.FindAllString(diag, -1)
		count += len(results)

		iresults := ire.FindAllString(diag, -1)
		count += len(iresults)

	}

	return count

}

func part1() {
	timeInit := time.Now()

	input, err := os.ReadFile("inputs/exer4.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	result := string(input)

	totalLines := len(strings.Split(result, "\r\n"))
	lineSize := 0

	for _, line := range strings.Split(result, "\r\n") {
		lineSize = len(line)
		break
	}

	grid := make([][]string, totalLines)

	for i, line := range strings.Split(result, "\r\n") {
		grid[i] = make([]string, lineSize)
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}

	hCount := countHorizontal(&grid, "XMAS")
	vCount := countVertical(&grid, "XMAS")
	dCount := countDiagonal(&grid, "XMAS")

	totalCount := hCount + vCount + dCount

	fmt.Printf("Exercise4 - Result: %v\n", totalCount)
	fmt.Printf("Took: %v\n", time.Since(timeInit))
}

func Solution() {
	part1()
}
