package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	totalY := 0
	grid := make([][]string, 1000, 1000)
	coord := [2]int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		line = strings.TrimSuffix(line, "\n")

		for x, ch := range line {
			if ch == '^' {
				coord[0] = x
				coord[1] = totalY
				ch = '.'
			}
			grid[totalY] = append(grid[totalY], string(ch))
		}
		totalY++
	}

	grid = grid[:totalY]

	dirs := [][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	visited := map[[2]int]bool{}
	curDir := 0
	answer := 0

	for {
		_, exists := visited[coord]
		if !exists {
			answer++
		}
		visited[coord] = true

		x := coord[0] + dirs[curDir][0]
		y := coord[1] + dirs[curDir][1]
		if x < 0 || x >= len(grid[0]) ||
			y < 0 || y >= totalY {
			break
		}

		if grid[y][x] == "." {
			coord[0] = x
			coord[1] = y
		} else {
			curDir = (curDir + 1) % 4
		}
	}

	fmt.Println("Answer:", answer)
}
