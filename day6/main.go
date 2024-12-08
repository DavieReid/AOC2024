package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	X int
	Y int
}

func (p *Position) String() {
	fmt.Printf("[%v][%v]\n", p.X, p.X)
}

type Guard struct {
	CurrentPosition Position
	Direction       string
	Gone            bool
	Movements       int
}

var visited = make(map[Position]bool)

func (g *Guard) track() {
	g.Movements++
	visited[g.CurrentPosition] = true
	fmt.Println("Guard Movements so far", g.Movements)
}

func (g *Guard) moveUp() {
	g.CurrentPosition.X = g.CurrentPosition.X - 1
	fmt.Println("Moving Up", g.CurrentPosition)
}

func (g *Guard) moveDown() {
	g.CurrentPosition.X = g.CurrentPosition.X + 1
	fmt.Println("Moving Down", g.CurrentPosition)
}

func (g *Guard) moveLeft() {
	g.CurrentPosition.Y = g.CurrentPosition.Y - 1
	fmt.Println("Moving Left", g.CurrentPosition)
}

func (g *Guard) moveRight() {
	g.CurrentPosition.Y = g.CurrentPosition.Y + 1
	fmt.Println("Moving Right", g.CurrentPosition)
}

func (g *Guard) turn() {
	switch g.Direction {
	case "up":
		g.Direction = "right"
	case "down":
		g.Direction = "left"
	case "left":
		g.Direction = "up"
	case "right":
		g.Direction = "down"
	}
	fmt.Println("Guard Turned", g.Direction)
}

func (g *Guard) isBlocked(grid [][]rune) bool {
	blocked := false
	switch g.Direction {
	case "up":
		if string(grid[g.CurrentPosition.X-1][g.CurrentPosition.Y]) == "#" {
			fmt.Printf("Blocked at [%v][%v]\n", g.CurrentPosition.X-1, g.CurrentPosition.Y)
			blocked = true
		}
	case "down":
		if string(grid[g.CurrentPosition.X+1][g.CurrentPosition.Y]) == "#" {
			fmt.Printf("Blocked at [%v][%v]\n", g.CurrentPosition.X+1, g.CurrentPosition.Y)
			blocked = true
		}
	case "left":
		if string(grid[g.CurrentPosition.X][g.CurrentPosition.Y-1]) == "#" {
			fmt.Printf("Blocked at [%v][%v]\n", g.CurrentPosition.X, g.CurrentPosition.Y-1)
			blocked = true
		}
	case "right":
		if string(grid[g.CurrentPosition.X][g.CurrentPosition.Y+1]) == "#" {
			fmt.Printf("Blocked at [%v][%v]\n", g.CurrentPosition.X, g.CurrentPosition.Y+1)
			blocked = true
		}
	}
	if blocked {
		fmt.Printf("CURRENT POSITION: [%v][%v]\n", g.CurrentPosition.X, g.CurrentPosition.Y)
	}
	return blocked
}

func (g *Guard) hasGone(rowCount, columnCount int) bool {
	gone := false
	switch g.Direction {
	case "up":
		if g.CurrentPosition.X-1 < 0 {
			fmt.Println("Guard has gone from top of map")
			gone = true
		}
	case "down":
		if g.CurrentPosition.X+1 > rowCount-1 {
			fmt.Println("Guard has gone from bottom of map")
			gone = true
		}
	case "left":
		if g.CurrentPosition.Y-1 < 0 {
			fmt.Println("Guard has gone from left of map")
			gone = true
		}
	case "right":
		if g.CurrentPosition.Y+1 > columnCount-1 {
			fmt.Println("Guard has gone from left of map")
			gone = true
		}
	}

	return gone
}

func main() {
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
	guard := &Guard{Movements: 0}

	for rowIndex, row := range grid {
		for columnIndex := range row {
			letter := string(grid[rowIndex][columnIndex])
			if letter == "^" {
				guard.CurrentPosition.X = rowIndex
				guard.CurrentPosition.Y = columnIndex
				guard.Direction = "up"
			}
		}
	}

	for !guard.Gone {
		if guard.hasGone(rowCount, columnCount) {
			break
		}

		if guard.isBlocked(grid) {
			guard.turn()
			continue
		}

		switch guard.Direction {
		case "up":
			guard.moveUp()
		case "right":
			guard.moveRight()
		case "left":
			guard.moveLeft()
		case "down":
			guard.moveDown()
		}
		guard.track()
	}

	fmt.Printf("%v distinct positions\n", len(visited))

}
