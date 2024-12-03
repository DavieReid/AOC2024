package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DataList struct {
	List []int
}

func (dl *DataList) AddItem(str string) {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	dl.Append(val)
}

func (dl *DataList) Append(item int) {
	dl.List = append(dl.List, item)
}

func (dl *DataList) Sort() {
	sort.Ints(dl.List)
}

func CalculateSimilarity(left *DataList, right *DataList) int {
	score := 0
	for _, val := range left.List {
		appearances := 0
		for j := range right.List {
			if val == right.List[j] {
				appearances += 1
			}
		}
		score += val * appearances
	}
	return score
}

func main() {
	fmt.Println("Advent of Code 2024 - day 1 part 2")
	left := &DataList{}
	right := &DataList{}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format. Expected 2 numbers on a line separated by a space:", line)
			continue
		}

		left.AddItem(parts[0])
		right.AddItem(parts[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	score := CalculateSimilarity(left, right)
	fmt.Println("Similarity is: ", score)
}
