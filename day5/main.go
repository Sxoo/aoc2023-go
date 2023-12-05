// finished part1 2h±
// finished part2 6min

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
//go:embed test.txt
var testInput string

type Range struct {
	startKey   int
	endKey     int
	offset     int
	rangeOf	   int
}

type SeedRange struct {
	startKey   int
	rangeOf	   int
}

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
	seeds, ranges := parseInput(input)
	min := math.MaxInt64

	for _, seed := range seeds {
		currentSeed := stringToInt(seed)
	
		for i := 1; i <= 7; i++ {
			for _, r := range ranges[i] {
				if currentSeed >= r.startKey && currentSeed <= r.endKey {
					currentSeed = currentSeed + r.offset
					break
				}
			}
			if (i == 7) {
				if (currentSeed < min) {
					min = currentSeed
				}
			}
		}
	}

	return min
}

func part2(input string) int {
	seedRanges, ranges := parseInput2(input)
	min := math.MaxInt64

	for _, seed := range seedRanges {
		rangeOf := seed.rangeOf
		for i :=0 ; i < rangeOf; i++ {
			currentSeed := seed.startKey + i
			for i := 1; i <= 7; i++ {
				for _, r := range ranges[i] {
					if currentSeed >= r.startKey && currentSeed <= r.endKey {
						currentSeed = currentSeed + r.offset
						break
					}
				}
				if (i == 7) {
					if (currentSeed < min) {
						min = currentSeed
					}
				}
			}
		}	
	}

	return min
}


func parseInput(input string) ([]string, map[int][]Range) {
	var seeds []string
	rangesMap := make(map[int][]Range)

	var currentRanges []Range
	var index = 0

	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds = strings.Fields(line)[1:]
		} else if strings.HasSuffix(line, "map:") {
			if len(currentRanges) > 0 {
				rangesMap[index] = currentRanges
			}

			currentRanges = []Range{}
			index++

			continue
		} else {
			numbers := strings.Fields(line)

			startingValue, _ := strconv.Atoi(numbers[0])
			startingKey, _ := strconv.Atoi(numbers[1])
			rangeOf, _ := strconv.Atoi(numbers[2])

			offset := startingValue - startingKey
			endKey := startingKey + rangeOf - 1

			currentRanges = append(currentRanges, Range{startingKey,endKey, offset, rangeOf})
		}
	}

	if len(currentRanges) > 0 {
		rangesMap[index] = currentRanges
	}
	
	return seeds, rangesMap
}

func parseInput2(input string) ([]SeedRange, map[int][]Range) {
	var seedsRanges []SeedRange
	rangesMap := make(map[int][]Range)

	var currentRanges []Range
	var index = 0

	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds := strings.Fields(line)[1:]
			for i := 0; i < len(seeds); i = i+2 {
				startKey := stringToInt(seeds[i])
				rangeOf := stringToInt(seeds[i + 1])
				seedsRanges = append(seedsRanges, SeedRange{startKey, rangeOf})
			}
		} else if strings.HasSuffix(line, "map:") {
			if len(currentRanges) > 0 {
				rangesMap[index] = currentRanges
			}

			currentRanges = []Range{}
			index++

			continue
		} else {
			numbers := strings.Fields(line)

			startingValue, _ := strconv.Atoi(numbers[0])
			startingKey, _ := strconv.Atoi(numbers[1])
			rangeOf, _ := strconv.Atoi(numbers[2])

			offset := startingValue - startingKey
			endKey := startingKey + rangeOf - 1

			currentRanges = append(currentRanges, Range{startingKey,endKey, offset, rangeOf})
		}
	}

	if len(currentRanges) > 0 {
		rangesMap[index] = currentRanges
	}

	return seedsRanges, rangesMap
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}