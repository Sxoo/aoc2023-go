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
	"unicode"
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
	var sum int

	parsed := parseInput(input)
	fmt.Println(parsed)

	for _, str := range parsed {
		var first rune
		var last rune
		var haveFirst bool

		for _, char := range str {
			if unicode.IsDigit(char) {
				if haveFirst {
					last = char
				}
			
				if !haveFirst {
					haveFirst = true
					first = char
					last = char
				}
			}
		}
		fmt.Printf("first: %c, last: %c\n", first, last)
		str := string(first) + string(last)

		sum = sum + stringToInt(str)
	}
	return sum
}

func convertNumberWords(input string) string {
		numberWords := map[string]string{
		"oneight": "18",
		"threeight": "38",
		"nineight": "98",
		"fiveight": "58",
		"eightwo": "82",
		"eighthree": "83",
		"twone": "21",
		"sevenine": "79",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

    re := regexp.MustCompile("(threeight|nineight|eighthree|oneight|eightwo|fiveight|twone|sevenine|one|two|six|eight|five|three|four|seven|nine|)")
	converted := re.ReplaceAllStringFunc(input, func(match string) string {
		if val, ok := numberWords[match]; ok {
			return val
		}
		return match
	})

	return converted
}



func part2(input string) int {
	var sum int

	parsed := parseInput(input)

	for _, str := range parsed {
		var first rune
		var last rune
		var haveFirst = false

		convertedText := convertNumberWords(str)
		fmt.Printf("original text: %s\n convertedText: %s\n",str, convertedText)

		for _, char := range convertedText {
			if unicode.IsDigit(char) {
				if haveFirst {
					last = char
				}
			
				if !haveFirst {
					haveFirst = true
					first = char
					last = char
				}
			}
		}
	
		str := string(first) + string(last)
		sum = sum + stringToInt(str)
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