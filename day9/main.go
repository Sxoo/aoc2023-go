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

func sumOfSequence(sequence []int) int {
	var sum int
	for _, num := range sequence {
		sum += num
	}

	return sum
}



func part1(input string) int {
	parsed := parseInput(input)
	var sequenceAnswer int

	for _, line := range parsed {
		currentLine := line[0]
		for sumOfSequence(currentLine) != 0 {
			var newSequence []int
			for i := 1; i < len(currentLine); i++ {
				newSequence = append(newSequence, currentLine[i] - currentLine[i-1])
			}
		
			line = append(line, newSequence)
			currentLine = newSequence
		}

		currN := 0
		for i := len(line) - 1; i >= 0; i-- {
			currN = currN + line[i][len(line[i]) - 1]
		}

		sequenceAnswer += currN
	}


	return sequenceAnswer
}

func part2(input string) int {
	parsed := parseInput(input)
	var sequenceAnswer int

	for _, line := range parsed {
		currentLine := line[0]
		for sumOfSequence(currentLine) != 0 {
			var newSequence []int
			for i := 1; i < len(currentLine); i++ {
				newSequence = append(newSequence, currentLine[i] - currentLine[i-1])
			}
		
			line = append(line, newSequence)
			currentLine = newSequence
		}


		currN := 0
		for i := len(line) - 1; i >= 0; i-- {
			currN = (currN - line[i][0]) * -1
		}

		sequenceAnswer += currN
	}


	return sequenceAnswer
}

func parseInput(input string) (sequenceMap map[int][][]int) {
	sequenceMap = make(map[int][][]int)

	for index, line := range strings.Split(input, "\n") {
		var lineOfInts []int
		split := strings.Fields(line)[0:]
		for _, num := range split {
			lineOfInts = append(lineOfInts, stringToInt(num))
		}
		sequenceMap[index] = append(sequenceMap[index], lineOfInts)
	}

	return sequenceMap
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}