// finished part1 11min
// finished part2 2min

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
	time, distance := parseInput(input)

	var multiplied = 1

	for index, t := range time {
		var beat = 0
		var distanceToBeat = stringToInt(distance[index])
		for i := 1; i < stringToInt(t); i++ {
			var travelTime = stringToInt(t) - i
			var distanceTravelled = i * travelTime
			
			if distanceTravelled > distanceToBeat {
				beat++
				}
		}

		if (beat > 0) {
			multiplied = multiplied * beat
		}

	}


	return multiplied
}

func part2(input string) int {
	time, distance := parseInput2(input)

	var multiplied = 1
	var beat = 0
	var distanceToBeat = stringToInt(distance)

	for i := 1; i < stringToInt(time); i++ {
		var travelTime = stringToInt(time) - i
		var distanceTravelled = i * travelTime
		
		if distanceTravelled > distanceToBeat {
			beat++
			}
	}

	if (beat > 0) {
		multiplied = multiplied * beat
	}

	return multiplied
}

func parseInput(input string) (time []string, distance []string) {
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Time:") {
			time = strings.Fields(line)[1:]
		}
		if strings.HasPrefix(line, "Distance:") {
			distance = strings.Fields(line)[1:]
		}
	}
	return time, distance
}

func parseInput2(input string) (time string, distance string) {
	var times []string
	var distances []string

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Time:") {
			times = strings.Fields(line)[1:]
			for _, num := range times {
				time += num
			}
		}
		if strings.HasPrefix(line, "Distance:") {
			distances = strings.Fields(line)[1:]
			for _, num := range distances {
				distance += num
			}
		}
	}

	return time, distance
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}