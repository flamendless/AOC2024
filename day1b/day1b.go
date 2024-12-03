package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const N = 1000
	left := make([]int, 0, N)
	right := make([]int, 0, N)

	similarity := map[int]int{}

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

		tokens := strings.Split(line, "   ")
		l, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)

		similarity[r] += 1
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sum := 0
	for i := 0; i < N; i++ {
		s, _ := similarity[left[i]]
		sum += left[i] * s
	}

	fmt.Println("Answer:", sum)
}