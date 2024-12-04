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
			if grid[y][x] != "M" {
				continue
			}

			//_
			//> M * M
			//  * A *
			//  S * S
			if x+2 < len(grid[y]) && y+2 < rows {
				mr := grid[y][x+2] == "M"
				a := grid[y+1][x+1] == "A"
				sl := grid[y+2][x] == "S"
				sr := grid[y+2][x+2] == "S"
				if mr && a && sl && sr {
					answer++
				}
			}

			//> M * S
			//  * A *
			//  M * S
			if x+2 < len(grid[y]) && y+2 < rows {
				sr := grid[y][x+2] == "S"
				sr2 := grid[y+2][x+2] == "S"
				a := grid[y+1][x+1] == "A"
				ml := grid[y+2][x] == "M"
				if sr && sr2 && a && ml {
					answer++
				}
			}

			//  S * S
			//  * A *
			//> M * M
			if x+2 < len(grid[y]) && y-2 >= 0 {
				sl := grid[y-2][x] == "S"
				sr := grid[y-2][x+2] == "S"
				a := grid[y-1][x+1] == "A"
				mr := grid[y][x+2] == "M"
				if sl && sr && a && mr {
					answer++
				}
			}

			//  S * M <
			//  * A *
			//  S * M
			if x-2 >= 0 && y+2 < len(grid[y]) {
				sl := grid[y][x-2] == "S"
				a := grid[y+1][x-1] == "A"
				sl2 := grid[y+2][x-2] == "S"
				mr := grid[y+2][x] == "M"
				if sl && a && sl2 && mr {
					answer++
				}
			}
		}
	}

	fmt.Println("Answer:", answer)
}
