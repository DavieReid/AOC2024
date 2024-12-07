package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OrderRule struct {
	A int
	B int
}

type Update struct {
	Entries  []int
	Valid    bool
	Modified bool
	Middle   int
}

func findValues(rules []OrderRule, field string, predicate func(OrderRule) bool) []int {
	var values []int
	for _, rule := range rules {
		if predicate(rule) {
			switch field {
			case "A":
				values = append(values, rule.A)
			case "B":
				values = append(values, rule.B)
			default:
				fmt.Println("Invalid field specified. Use 'A' or 'B'.")
				return nil
			}
		}
	}
	return values
}

func intersection(slice1, slice2 []int) []int {
	elementMap := make(map[int]bool)
	var result []int

	// Populate the map with elements from the first slice
	for _, num := range slice1 {
		elementMap[num] = true
	}

	// Check for common elements in the second slice
	for _, num := range slice2 {
		if elementMap[num] {
			result = append(result, num)
			// Remove the element from the map to avoid duplicates in the result
			delete(elementMap, num)
		}
	}

	return result
}

func swap(update *Update, index1, index2 int) {
	// Swap the elements at the given indices
	update.Modified = true
	update.Entries[index1], update.Entries[index2] = update.Entries[index2], update.Entries[index1]
}

func main() {
	orderingRules := []OrderRule{}
	updates := []Update{}

	file, err := os.Open("rules.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "|")

		a, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		b, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}

		rule := OrderRule{A: a, B: b}

		orderingRules = append(orderingRules, rule)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer input.Close()
	scanner = bufio.NewScanner(input)
	beforeRules := make(map[int][]int)
	afterRules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		entry := Update{Valid: true}

		for _, val := range parts {
			updateNum, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}

			beforeRules[updateNum] = findValues(orderingRules, "B", func(or OrderRule) bool { return or.A == updateNum })
			afterRules[updateNum] = findValues(orderingRules, "A", func(or OrderRule) bool { return or.B == updateNum })
			entry.Entries = append(entry.Entries, updateNum)

		}
		updates = append(updates, entry)

	}

	invalidUpdates := []Update{}
	for _, update := range updates {
		isValid := true
		for i := 0; i < len(update.Entries); i++ {
			fmt.Println("i", i, update.Entries)
			updateVal := update.Entries[i]
			beforeElems := update.Entries[:i]
			afterElems := update.Entries[i+1:]
			isFirst := len(beforeElems) == 0
			isLast := len(afterElems) == 0

			// fmt.Printf("Checking %v... must be before %v and after %v\n", updateVal, beforeRules[updateVal], afterRules[updateVal])

			// fmt.Println("Elements before:", beforeElems)
			// fmt.Println("Elements after: ", afterElems)

			// if any element that is in the before rules for this updateVal appears in the afterElems = FAIL
			beforeCheck := intersection(beforeElems, beforeRules[updateVal])
			// if any element that is in the after rules appears in the beforeElems = FAIL
			isAfter := intersection(afterElems, afterRules[updateVal])

			if isFirst && len(beforeCheck) > 0 {
				fmt.Println("FIRST WE have a failure", updateVal)
				isValid = false
				swap(&update, i, i+1)
			}

			if isLast && len(isAfter) > 0 {
				fmt.Println("LAST WE have a failure", updateVal)
				swap(&update, i, i-1)
				fmt.Println(update.Entries)
				isValid = false
			}

			if len(beforeCheck) > 0 {
				fmt.Println("Broken Before rule", updateVal, beforeCheck)
				isValid = false
				swap(&update, i, i-1)
			}

			if len(isAfter) > 0 {
				fmt.Println("Broken After rule", updateVal, isAfter)
				isValid = false
				swap(&update, i, i+1)
			}

			if !isValid && isLast {
				i = 0 // reset the loop
				isValid = true
			}
		}

		if isValid {
			invalidUpdates = append(invalidUpdates, update)
		}

	}
	total := 0
	for _, val := range invalidUpdates {
		if val.Modified {
			total += val.Entries[len(val.Entries)/2]
		}
	}
	fmt.Println("Total Invalid Updates: ", invalidUpdates)
	fmt.Println("Sum of Middle Entries", total)

}
