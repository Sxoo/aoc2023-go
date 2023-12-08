// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
//go:embed test.txt
var testInput string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	testInput = strings.TrimRight(testInput, "\n")
	if len(testInput) == 0 {
		panic("empty test.txt file")
	}
}

func main() {
	var part int
	var test bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&test, "test", false, "run with test.txt inputs?")
	flag.Parse()
	fmt.Println("Running part", part, ", test inputs = ", test)

  if test {
		input = testInput
	}

	var ans int
	switch part {
	case 1:
		ans = part1(input)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	direction, nodeMap := parseInput(input)

	var steps = 0
	var key = "AAA"
	
	for key != "ZZZ" {
		for _, char := range direction {
			if char == 'L' {
				key = nodeMap[key]["L"]
				steps++
				if key == "ZZZ" {
					break;
				}
			} else if char == 'R' {
				key = nodeMap[key]["R"]
				steps++
				if key == "ZZZ" {
					break;
				}
			}
		}
	}

	return steps
}

// src: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func part2(input string) int {
	direction, nodeMap, startingKeys := parseInput2(input)
	var finalSteps []int

	for _, key := range startingKeys {
		var steps = 0
		for string(key[2]) != "Z" {
			for _, char := range direction {
				if char == 'L' {
					key = nodeMap[key]["L"]
					steps++
					if string(key[2]) == "Z" {
						break;
					}
				} else if char == 'R' {
					key = nodeMap[key]["R"]
					steps++
					if string(key[2]) == "Z" {
						break;
					}
				}
			}
		}
		finalSteps = append(finalSteps, steps)
	}

	ans := LCM(finalSteps[0], finalSteps[1], finalSteps[2], finalSteps[3], finalSteps[4], finalSteps[5])

	return ans
}


func parseInput(input string) (directions string, nodeMap map[string]map[string]string) {
	parts := strings.Split(input, "\n\n")
	directions = parts[0]
	nodeMap = make(map[string]map[string]string)

	for _, line := range strings.Split(parts[1], "\n") {
		lineParts := strings.Split(line, " = ")
		
		trimLeft := strings.Replace(lineParts[1], "(", "", -1)
		trimRight := strings.Replace(trimLeft, ")", "", -1)
		
		leftRightValue := strings.Split(trimRight, ", ")

		nodeMap[lineParts[0]] = map[string]string{
		"L": leftRightValue[0],
		"R": leftRightValue[1],
		}
	}

	return directions, nodeMap
}

func parseInput2(input string) (directions string, nodeMap map[string]map[string]string, startingKeys []string) {
	parts := strings.Split(input, "\n\n")
	directions = parts[0]
	nodeMap = make(map[string]map[string]string)

	for _, line := range strings.Split(parts[1], "\n") {
		lineParts := strings.Split(line, " = ")
		if (lineParts[0][2] == 'A') {
			startingKeys = append(startingKeys, lineParts[0])
		}
		trimLeft := strings.Replace(lineParts[1], "(", "", -1)
		trimRight := strings.Replace(trimLeft, ")", "", -1)
		
		leftRightValue := strings.Split(trimRight, ", ")

		nodeMap[lineParts[0]] = map[string]string{
		"L": leftRightValue[0],
		"R": leftRightValue[1],
		}
	}

	return directions, nodeMap, startingKeys
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}