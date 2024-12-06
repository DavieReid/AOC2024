package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2024 - day 4")
	forwards := "XMAS"
	backwards := "SAMX"
	wordLength := len(forwards)

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
			word := make([]rune, wordLength)
			if letter == 'X' {
				// check row left ==> right
				if columnIndex+3 <= columnCount {
					substr := string(row[columnIndex : columnIndex+wordLength])

					if substr == forwards {
						fmt.Println("FOUND FORWARDS ON THE ROW", rowIndex, columnIndex)
						count += 1
					}
				}

				// check column down
				for i := 0; i <= 3; i++ {
					if rowIndex+i < rowCount {
						word[i] = grid[rowIndex+i][columnIndex]
					}
				}

				if string(word) == forwards {
					fmt.Println("FOUND FORWARDS STRAIGHT DOWN", rowIndex, columnIndex)
					count += 1
				}

				// check column diagonal left and down
				word = make([]rune, wordLength)
				for i := 0; i <= 3; i++ {
					if rowIndex+i < rowCount && columnIndex+i < columnCount {
						word[i] = grid[rowIndex+i][columnIndex+i]
					}
				}

				if string(word) == forwards {
					fmt.Println("FOUND FORWARDS ON THE DIAGONAL RIGHT AND DOWN", rowIndex, columnIndex)
					count += 1
				}

				// check column diagonal left and down
				word = make([]rune, wordLength)
				for i := 0; i <= 3; i++ {
					if rowIndex+i < rowCount && columnIndex-i >= 0 {
						word[i] = grid[rowIndex+i][columnIndex-i]
					}
				}
				if string(word) == forwards {
					fmt.Println("FOUND FORWARDS ON THE DIAGONAL LEFT AND DOWN", rowIndex, columnIndex)
					count += 1
				}

			}

			if letter == 'S' {
				// check row left ==> right
				if columnIndex+3 <= columnCount {
					substr := string(row[columnIndex : columnIndex+wordLength])
					if substr == backwards {
						fmt.Println("FOUND BACKWARDS ON THE ROW", rowIndex, columnIndex)
						count += 1
					}
				}

				// check column down
				for i := 0; i <= 3; i++ {
					if rowIndex+i < rowCount {
						word[i] = grid[rowIndex+i][columnIndex]
					}
				}

				if string(word) == backwards {
					count += 1
					fmt.Println("FOUND BACKWARDS STRAIGHT DOWN", rowIndex, columnIndex)
				}

				// check column diagonal right and down
				word = make([]rune, wordLength)
				for i := 0; i <= 3; i++ {
					if rowIndex+i < rowCount && columnIndex+i < columnCount {
						word[i] = grid[rowIndex+i][columnIndex+i]
					}
				}

				if string(word) == backwards {
					fmt.Println("FOUND BACKWARDS ON THE DIAGONAL RIGHT AND DOWN", rowIndex, columnIndex)
					count += 1
				}

				// check column diagonal left and down
				word = make([]rune, wordLength)
				for i := 0; i <= 3; i++ {
					if rowIndex+i < rowCount && columnIndex-i >= 0 {
						word[i] = grid[rowIndex+i][columnIndex-i]
					}
				}
				if string(word) == backwards {
					fmt.Println("FOUND BACKWARDS ON THE DIAGONAL LEFT AND DOWN", rowIndex, columnIndex)
					count += 1
				}

			}

		}
	}

	fmt.Printf("XMAS found %v times.\n", count)

}
