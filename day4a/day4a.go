package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := make([][]string, 256)

	rows := 0
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

		grid[rows] = make([]string, 0, len(line))

		for _, ch := range line {
			grid[rows] = append(grid[rows], string(ch))
		}

		rows++
	}

	answer := 0
	for y := range rows {
		for x := range len(grid[y]) {
			if grid[y][x] != "X" {
				continue
			}

			//left to right
			if x+3 < len(grid[y]) {
				m := grid[y][x+1] == "M"
				a := grid[y][x+2] == "A"
				s := grid[y][x+3] == "S"
				if m && a && s {
					answer++
				}
			}

			//right to left
			if x-3 >= 0 {
				m := grid[y][x-1] == "M"
				a := grid[y][x-2] == "A"
				s := grid[y][x-3] == "S"
				if m && a && s {
					answer++
				}
			}

			//top to bottom
			if y+3 < rows {
				m := grid[y+1][x] == "M"
				a := grid[y+2][x] == "A"
				s := grid[y+3][x] == "S"
				if m && a && s {
					answer++
				}
			}

			//bottom to top
			if y-3 >= 0 {
				m := grid[y-1][x] == "M"
				a := grid[y-2][x] == "A"
				s := grid[y-3][x] == "S"
				if m && a && s {
					answer++
				}
			}

			//diagonal top left to bottom right
			if y+3 < rows && x+3 < len(grid[y]) {
				m := grid[y+1][x+1] == "M"
				a := grid[y+2][x+2] == "A"
				s := grid[y+3][x+3] == "S"
				if m && a && s {
					answer++
				}
			}

			//diagonal top right to bottom left
			if y+3 < rows && x-3 >= 0 {
				m := grid[y+1][x-1] == "M"
				a := grid[y+2][x-2] == "A"
				s := grid[y+3][x-3] == "S"
				if m && a && s {
					answer++
				}
			}

			//diagonal bottom left to top right
			if y-3 >= 0 && x+3 < len(grid[y]) {
				m := grid[y-1][x+1] == "M"
				a := grid[y-2][x+2] == "A"
				s := grid[y-3][x+3] == "S"
				if m && a && s {
					answer++
				}
			}

			//diagonal bottom right to top left
			if y-3 >= 0 && x-3 >= 0 {
				m := grid[y-1][x-1] == "M"
				a := grid[y-2][x-2] == "A"
				s := grid[y-3][x-3] == "S"
				if m && a && s {
					answer++
				}
			}
		}
	}

	fmt.Println("Answer:", answer)
}
