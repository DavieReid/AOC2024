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

type Obstacle struct {
	Position           Position
	DamagedLeftCount   int
	DamagedRightCount  int
	DamagedTopCount    int
	DamagedBottomCount int
}

var visited = make(map[Position]bool)

func (g *Guard) track() {
	g.Movements++
	visited[g.CurrentPosition] = true
	// fmt.Println("Guard Movements so far", g.Movements)
}

func (g *Guard) moveUp() {
	g.CurrentPosition.X = g.CurrentPosition.X - 1
	//fmt.Println("Moving Up", g.CurrentPosition)
}

func (g *Guard) moveDown() {
	g.CurrentPosition.X = g.CurrentPosition.X + 1
	//fmt.Println("Moving Down", g.CurrentPosition)
}

func (g *Guard) moveLeft() {
	g.CurrentPosition.Y = g.CurrentPosition.Y - 1
	//fmt.Println("Moving Left", g.CurrentPosition)
}

func (g *Guard) moveRight() {
	g.CurrentPosition.Y = g.CurrentPosition.Y + 1
	//fmt.Println("Moving Right", g.CurrentPosition)
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
	// fmt.Println("Guard Turned", g.Direction)
}

func (g *Guard) isBlocked(grid [][]rune) (bool, bool) {
	var next Position
	switch g.Direction {
	case "up":
		next = Position{X: g.CurrentPosition.X - 1, Y: g.CurrentPosition.Y}
	case "down":
		next = Position{X: g.CurrentPosition.X + 1, Y: g.CurrentPosition.Y}
	case "left":
		next = Position{X: g.CurrentPosition.X, Y: g.CurrentPosition.Y - 1}
	case "right":
		next = Position{X: g.CurrentPosition.X, Y: g.CurrentPosition.Y + 1}
	}
	if string(grid[next.X][next.Y]) == "#" {
		//fmt.Printf("Blockage on [%v][%v]\n", next.X, next.Y)
		//fmt.Printf("CURRENT POSITION: [%v][%v]\n", g.CurrentPosition.X, g.CurrentPosition.Y)
		return true, false
	}

	if string(grid[next.X][next.Y]) == "O" {
		//fmt.Printf("CURRENT POSITION: [%v][%v]\n", g.CurrentPosition.X, g.CurrentPosition.Y)
		//fmt.Printf("!!!!!!!!OBSTACLE!!!!!!! at [%v][%v]\n", next.X, next.Y)
		return false, true
	}
	return false, false
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

func run(guard Guard, grid [][]rune, rowCount int, columnCount int, obstacle Obstacle) bool {
	//fmt.Println("Obstacle placed at:", obstacle.Position.X, obstacle.Position.Y)
	//fmt.Println("Guard starts at", guard.CurrentPosition.X, guard.CurrentPosition.Y)

	for {
		if guard.hasGone(rowCount, columnCount) {
			return false
		}
		guard.track()
		blocked, blockedByObstacle := guard.isBlocked(grid)

		if blockedByObstacle {
			switch guard.Direction {
			case "up":
				obstacle.DamagedBottomCount++
			case "down":
				obstacle.DamagedTopCount++
			case "left":
				obstacle.DamagedRightCount++
			case "right":
				obstacle.DamagedLeftCount++
			}

			if obstacle.DamagedLeftCount > 1 || obstacle.DamagedTopCount > 1 || obstacle.DamagedRightCount > 1 || obstacle.DamagedBottomCount > 1 {
				return true
			}
			guard.turn()
			continue
		}

		if blocked {
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
	}
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

	loopsDetected := 0
	guard := Guard{Movements: 0, Gone: false}
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

	for rowIndex, row := range grid {
		for columnIndex := range row {
			letter := string(grid[rowIndex][columnIndex])
			var obstacle Obstacle
			if letter != "#" && letter != "^" {
				grid[rowIndex][columnIndex] = 'O'
				obstacle = Obstacle{Position: Position{X: rowIndex, Y: columnIndex}, DamagedBottomCount: 0, DamagedLeftCount: 0, DamagedRightCount: 0, DamagedTopCount: 0}
			}

			loopDetected := run(guard, grid, rowCount, columnCount, obstacle)
			fmt.Println(loopsDetected)
			if loopDetected {
				loopsDetected++
				fmt.Println("Loop")
			}

			//reset
			if grid[rowIndex][columnIndex] == 'O' {
				grid[rowIndex][columnIndex] = '.'
			}
		}
	}

	fmt.Println("LOOPS: ", loopsDetected)

}
