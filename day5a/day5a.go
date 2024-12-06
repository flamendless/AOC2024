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

	edges := map[int][]int{}

	processLastSection := false
	rules := make([]string, 0, 1000)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if len(strings.TrimSpace(line)) == 0 {
			processLastSection = true
			continue
		}

		line = strings.TrimSuffix(line, "\n")

		if !processLastSection {
			tokens := strings.Split(line, "|")
			l, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			r, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			edges[l] = append(edges[l], r)
		} else {
			rules = append(rules, line)
		}
	}

	fmt.Println(edges)
	fmt.Println(rules)

	for _, rule := range rules {
		set := map[int]bool{}
		valid := true
		tokens := strings.Split(rule, ",")
		for _, token := range tokens {
			x, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			for _, r := range edges[x] {
				_, exists := set[r]
				if exists {
					valid = false
					break
				}
			}
			set[x] = true
		}
		if valid {
			mid := tokens[len(tokens)/2]
			fmt.Println(mid)
			midN, err := strconv.Atoi(mid)
			if err != nil {
				panic(err)
			}
			answer += midN
		}
	}

	fmt.Println("Answer:", answer)
}
