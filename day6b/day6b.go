package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirs = [][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}
var grid = make([][]string, 1000, 1000)
var start = [2]int{}
var visited map[int]bool

func main() {
	totalY := 0

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
				start[0] = x
				start[1] = totalY
				ch = '.'
			}
			grid[totalY] = append(grid[totalY], string(ch))
		}
		totalY++
	}

	grid = grid[:totalY]
	visited = make(map[int]bool, len(grid) * len(grid[0]) * 4)

	answer := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != "." {
				continue
			}

			if start[0] != x || start[1] != y {
				grid[y][x] = "#"
				if getAnswer() {
					answer++
				}
				grid[y][x] = "."
			}
		}
	}

	fmt.Println("Answer:", answer)
}

func getAnswer() bool {
	coord := start
	curDir := 0
	for k := range visited {
		delete(visited, k)
	}

	for {
		key := (coord[1]*len(grid[0])+coord[0])*4 + curDir
		if val, exists := visited[key]; exists && val {
			return true
		}
		visited[key] = true

		x := coord[0] + dirs[curDir][0]
		y := coord[1] + dirs[curDir][1]
		if x < 0 || x >= len(grid[0]) ||
			y < 0 || y >= len(grid) {
			return false
		}

		if grid[y][x] == "." {
			coord[0] = x
			coord[1] = y
		} else {
			curDir = (curDir + 1) % 4
		}
	}
}
