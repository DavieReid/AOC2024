package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2024 - day 4")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	rowCount := len(grid)
	columnCount := len(grid[0])
	fmt.Printf("Grid size is %v rows by %v columns\n", rowCount, columnCount)
	count := 0

	for rowIndex, row := range grid {
		for columnIndex := range row {
			letter := grid[rowIndex][columnIndex]
			if (rowIndex > 0 && rowIndex+1 < rowCount) && (columnIndex > 0 && columnIndex+1 < columnCount) {

				if letter == 'A' {

					topLeft := grid[rowIndex-1][columnIndex-1]
					topRight := grid[rowIndex-1][columnIndex+1]
					bottomLeft := grid[rowIndex+1][columnIndex-1]
					bottomRight := grid[rowIndex+1][columnIndex+1]

					if topLeft == 'M' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'S' {
						count += 1
					}

					if topLeft == 'S' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'M' {
						count += 1
					}

					if topLeft == 'S' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'M' {
						count += 1
					}

					if topLeft == 'M' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'S' {
						count += 1
					}

				}
			}

		}
	}

	fmt.Printf("X-MAS found %v times.\n", count)

}
