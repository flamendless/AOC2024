package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	edges := map[int][]int{}
	processLastSection := false
	rulesLines := make([]string, 0, 1000)
	rules := make([][]int, 1000, 1000)
	rulesN := 0

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
			rulesLines = append(rulesLines, line)
			tokens := strings.Split(line, ",")
			rules[rulesN] = make([]int, 0, len(tokens))
			for _, token := range tokens {
				x, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				rules[rulesN] = append(rules[rulesN], x)
			}
			rulesN++
		}
	}

	answer := 0
	for _, rule := range rules {
		if len(rule) == 0 {
			break
		}

		set := map[int]bool{}
		valid := true

		in := map[int]int{}
		found := map[int]bool{}
		for _, n := range rule {
			found[n] = true
		}

		for _, n := range rule {
			for _, r := range edges[n] {
				if _, exists := found[r]; exists {
					in[r] += 1
				}

				if _, exists := set[r]; exists {
					valid = false
				}
			}
			set[n] = true
		}

		if !valid {
			starting := []int{}
			for _, n := range rule {
				if v, _ := in[n]; v == 0 {
					starting = append(starting, n)
				}
			}

			order := []int{}
			for _, st := range starting {
				sort(edges, st, &order, in, found)
			}

			mid := order[len(order)/2]
			answer += mid
		}
	}

	fmt.Println("Answer:", answer)
}

func sort(
	edges map[int][]int,
	x int,
	order *[]int,
	in map[int]int,
	found map[int]bool,
) {
	*order = append(*order, x)
	for _, edge := range edges[x] {
		_, exists := found[edge]
		if exists {
			in[edge] -= 1
			v, _ := in[edge]
			if v == 0 {
				sort(edges, edge, order, in, found)
			}
		}
	}
}
