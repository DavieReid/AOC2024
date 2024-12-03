package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	List []int
	Safe bool
	Type string
}

func (r *Report) AddItem(str string) {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	if !r.Safe {
		fmt.Println("Report already marked as unsafe")
		return
	}

	length := len(r.List)

	if length == 0 {
		r.Append(val)
		return
	}

	lastItem := r.List[length-1]
	diff := lastItem - val

	if diff > 0 && r.Type == "" {
		r.Type = "DECREASING"
	}

	if diff < 0 && r.Type == "" {
		r.Type = "INCREASING"
	}

	if r.Type == "INCREASING" {
		diff = diff * -1
	}

	r.Safe = diff >= 1 && diff <= 3
	//fmt.Println(r.List, r.Safe, r.Type, diff, val)
	r.Append(val)
}

func (r *Report) Append(item int) {
	r.List = append(r.List, item)
}

func main() {
	fmt.Println("Advent of Code 2024 - day 2")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	reports := []Report{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		report := Report{
			Safe: true,
		}

		for _, val := range parts {
			report.AddItem(val)
		}

		reports = append(reports, report)

	}

	totalSafe := 0

	for _, r := range reports {
		if r.Safe {
			totalSafe += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total number of Safe Reports: ", totalSafe)
}
