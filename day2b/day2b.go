package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkSafe(tokens []string) (bool, int) {
	isSafe := true
	dir := 0
	unsafeIdx := 0

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
			unsafeIdx = i
			break
		}

		if l < r {
			if dir == -1 {
				isSafe = false
				unsafeIdx = i
				break
			}
			dir = 1
		}

		if l > r {
			if dir == 1 {
				isSafe = false
				unsafeIdx = i
				break
			}
			dir = -1
		}
	}

	return isSafe, unsafeIdx
}

func checkSafeAgain(tokens []string, i int) bool {
	newTokens := slices.Clone(tokens)
	newTokens = slices.Delete(newTokens, i, i+1)
	nowSafe, _ := checkSafe(newTokens)
	return nowSafe
}

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

		isSafe, unsafeIdx := checkSafe(tokens)
		isSafe = checkSafeAgain(tokens, 0)
		if !isSafe {
			isSafe = checkSafeAgain(tokens, unsafeIdx)
			if isSafe {
				answer += 1
				continue
			}
			isSafe = checkSafeAgain(tokens, unsafeIdx+1)
			if isSafe {
				answer += 1
				continue
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
