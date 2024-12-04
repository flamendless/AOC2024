package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getProduct(validLine string) (int, error) {
	line := strings.TrimPrefix(validLine, "mul(")
	line = strings.TrimSuffix(line, ")")
	tokens := strings.Split(line, ",")
	if len(tokens) != 2 {
		return 0, errors.New("Invalid tokens")
	}

	l, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, err
	}

	r, err := strconv.Atoi(tokens[1])
	if err != nil {
		return 0, err
	}

	return l * r, nil
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

		for i := 0; i < len(line); i++ {
			if line[i] != 'm' {
				continue
			}

			if i+4 < len(line) {
				substr := line[i : i+4]
				if substr != "mul(" {
					continue
				}

				j := i + 4
				for {
					if line[j] == ')' {
						break
					}
					j++
				}

				validLine := line[i : j+1]
				prod, err := getProduct(validLine)
				if err != nil {
					continue
				}
				answer += prod
			}
		}
	}

	fmt.Println("Answer:", answer)
}
