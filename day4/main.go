// finished part1 27min
// finished part2 31min

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

func part1(input string) int {
	parsedMap := parseInput(input)
	var sum int

	for _, card := range parsedMap {
		var cardSum int
		var isFirst bool = true

		for _, num := range card.myNumbers {
			if card.winningMap[num] {
				if !isFirst {
					cardSum = (cardSum * 2)
					continue
				}
				if isFirst {
					cardSum += 1
					isFirst = false
				}
			}
		}
		sum += cardSum
	}

	return sum
}

func part2(input string) int {
	parsedMap := parseInput(input)
	cardRepeats := make(map[int]int)
	var sum int

	// TIL: In Go, the for range loop over a map is unordered, meaning it doesn't guarantee the order of iteration. 
	// Maps are unordered collections. 
	// If you specifically want to iterate over a map in a predictable order, you can't rely on the order in which the elements were added because maps don't maintain order.
	keys := make([]int, 0, len(parsedMap))
	for key := range parsedMap {
		keys = append(keys, key)
	}

	for _, key := range keys {
		var winners = 0

		for i := 0; i <= cardRepeats[key]; i++ {
			var cardWinners = 0
			
			for _, num := range parsedMap[key].myNumbers {
				if parsedMap[key].winningMap[num] {
					cardWinners += 1
					key := key + cardWinners
					cardRepeats[key] += 1
				}
			}
			winners += cardWinners
		}

		sum += winners
	}

	return sum + len(parsedMap)
}

func parseInput(input string) (parsedMap map[int]struct {
		winningMap map[string]bool
		myNumbers    []string
	}) {


	parsedMap = make(map[int]struct {
		winningMap map[string]bool
		myNumbers    []string
	})
	
	for index, line := range strings.Split(input, "\n") {
		winningMap := make(map[string]bool)
		removedCardNumber := strings.Split(line, ":")
		cardSplit := strings.Split(removedCardNumber[1], "|")
		re:=regexp.MustCompile(`\d+`)
		winningNumbers := re.FindAllString(cardSplit[0], -1)
		myNumbers := re.FindAllString(cardSplit[1], -1)

		for _, num := range winningNumbers {
			winningMap[num] = true
		}

		parsedMap[index] = struct{winningMap map[string]bool; myNumbers []string}{winningMap, myNumbers}

	}
	return parsedMap
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}