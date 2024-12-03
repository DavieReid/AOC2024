package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func extractNumbers(s string) [][2]int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	var results [][2]int
	matches := re.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		if len(match) == 3 {
			var x, y int
			fmt.Sscanf(match[1], "%d", &x)
			fmt.Sscanf(match[2], "%d", &y)
			results = append(results, [2]int{x, y})
		}
	}
	return results
}

func main() {
	fmt.Println("Advent of Code 2024 - day 3")

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	content := string(data)
	// input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	muls := extractNumbers(content)

	result := 0
	for _, val := range muls {
		result += val[0] * val[1]
	}

	fmt.Println("RESULT: ", result)
}
