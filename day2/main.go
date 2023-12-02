// started        ;
// finished part1 , 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

func splitAndExtractGame(str string) (gameId int, line string) {
	string := strings.Split(str, ":")
	re := regexp.MustCompile("[0-9]+")
	id := re.FindString(string[0])

	return stringToInt(id), string[1]
}

func Splitter(s string, splits string) []string {
    m := make(map[rune]int)
    for _, r := range splits {
        m[r] = 1
    }

    splitter := func(r rune) bool {
        return m[r] == 1
    }

    return strings.FieldsFunc(s, splitter)
}

func part1(input string) int {
	colorMax := map[string]int{
		"red": 12,
		"blue": 14,
		"green": 13,
	}

	parsed := parseInput(input)
	var sum int

	for _, str := range parsed {
		gameId, line := splitAndExtractGame(str)
		sum = sum + gameId
		cubes := Splitter(line, ";,")
		
		for _, cube := range cubes {
			trimmed := strings.Trim(cube, " ")
			cubeSplit := strings.Split(trimmed, " ")
			max := colorMax[cubeSplit[1]]
			if (max < stringToInt(cubeSplit[0])) {
				    fmt.Println("Game", gameId, "failed")
					sum = sum - gameId
					break
				}	
			}
		}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)

	var sum int

	for _, str := range parsed {
		colorMin := map[string]int{
		"red": 0,
		"blue": 0,
		"green": 0,
		}

		gameId, line := splitAndExtractGame(str)
		cubes := Splitter(line, ";,")
		
		for _, cube := range cubes {
			trimmed := strings.Trim(cube, " ")
			cubeSplit := strings.Split(trimmed, " ")
			min := colorMin[cubeSplit[1]]
			if (min < stringToInt(cubeSplit[0])) {
				fmt.Println("Game", gameId, "new min req", cubeSplit[0])
				colorMin[cubeSplit[1]] = stringToInt(cubeSplit[0])
			}	
			}
		var sumOfGame int
		sumOfGame = colorMin["red"] * colorMin["blue"] * colorMin["green"]

		fmt.Println("gameid: ", gameId, "sumOfGame: ", sumOfGame)

		sum = sumOfGame + sum
		}
	return sum
}

func parseInput(input string) (parsedInput []string) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, line)
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}