package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := 0

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
		tokens := strings.Split(line, " ")

		isSafe := true
		dir := 0

		for i := 0; i < len(tokens)-1; i++ {
			left := tokens[i]
			l, err := strconv.Atoi(left)
			if err != nil {
				panic(err)
			}

			right := tokens[i+1]
			r, err := strconv.Atoi(right)
			if err != nil {
				panic(err)
			}

			diff := abs(l - r)
			if diff < 1 || diff > 3 {
				isSafe = false
				break
			}

			if l < r {
				if dir == -1 {
					isSafe = false
					break
				}
				dir = 1
			}

			if l > r {
				if dir == 1 {
					isSafe = false
					break
				}
				dir = -1
			}
		}

		if isSafe {
			answer += 1
		}
	}

	fmt.Println("Answer:", answer)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
