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

	do := true
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

		N := len(line)
		for i := 0; i < N; i++ {
			if line[i] == 'd' {
				if i+4 < N {
					substrDo := line[i : i+4]
					if substrDo == "do()" {
						do = true
					}
				}

				if i+7 < N {
					substrDont := line[i : i+7]
					if substrDont == "don't()" {
						do = false
					}
				}
			} else if line[i] == 'm' {
				if i+4 < N {
					substr := line[i : i+4]
					if substr != "mul(" {
						continue
					}

					j := i + 1
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

					if do {
						answer += prod
					}
				}
			}
		}
	}

	fmt.Println("Answer:", answer)
}
