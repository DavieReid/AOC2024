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

func CalculateDistance(left *DataList, right *DataList) int {
	distance := make([]int, len(left.List))
	for i := range left.List {
		diff := (left.List[i] - right.List[i])
		if diff < 0 {
			diff = diff * -1
		}
		distance[i] = diff
	}

	sum := 0
	for _, value := range distance {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println("Advent of Code 2024 - day 1")
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

	left.Sort()
	right.Sort()

	distance := CalculateDistance(left, right)
	fmt.Println("Distance is: ", distance)
}
