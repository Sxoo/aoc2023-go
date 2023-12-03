// finished part1 in 2h
// finished part2 in 39mins

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

func getSymbol(symbolMap map[string]string, row int, col int) string {
	return symbolMap[strconv.Itoa(row)+"!"+fmt.Sprint(col)]
}

func part1(input string) int {
	symbolMap, numberMap:= parseInput(input)
	var sum int

	for index, number := range numberMap {
		rowsAndCols := strings.Split(index, "!")
		rowAsInt := stringToInt(rowsAndCols[0])
		wholeColAsString := rowsAndCols[1]
		colAsString := strings.Split(wholeColAsString, "?")
		numberStart := stringToInt(colAsString[0])
		numberEnd := stringToInt(colAsString[1])

		for i := numberStart; i <= numberEnd; i++ {
			topRowIndex := rowAsInt - 1
			bottomRowIndex := rowAsInt + 1
			leftColIndex := i - 1
			rightColIndex := i + 1

			if getSymbol(symbolMap, topRowIndex, i) != "" ||
			getSymbol(symbolMap, bottomRowIndex, i) != "" ||
			getSymbol(symbolMap, rowAsInt, leftColIndex) != "" ||
			getSymbol(symbolMap, rowAsInt, rightColIndex) != "" ||
			getSymbol(symbolMap, topRowIndex, rightColIndex) != "" ||
			getSymbol(symbolMap, topRowIndex, leftColIndex) != "" ||
			getSymbol(symbolMap, bottomRowIndex, rightColIndex) != "" ||
			getSymbol(symbolMap, bottomRowIndex, leftColIndex) != "" {
				sum = sum + number
				break;
			} else {
				fmt.Println("No neighbours for", number)
			}
		}
	}

	return sum
}

func part2(input string) int {
	symbolMap, numberMap := parseInputPart2(input)
	var sum int

	for index := range symbolMap {
		var results []int

		split := strings.Split(index, "!")

		rowAsInt := stringToInt(split[0])
		colAsInt := stringToInt(split[1])

		topRowIndex := rowAsInt - 1
		bottomRowIndex := rowAsInt + 1
		leftColIndex := colAsInt - 1
		rightColIndex := colAsInt + 1


		val1, exists1 := numberMap[strconv.Itoa(topRowIndex)+"!"+fmt.Sprint(colAsInt)]
		if exists1 {
			results = append(results, val1)
		}
		val2, exists2 := numberMap[strconv.Itoa(bottomRowIndex)+"!"+fmt.Sprint(colAsInt)] 
		if exists2 {
			results = append(results, val2)
		}
		val3, exists3 := numberMap[strconv.Itoa(rowAsInt)+"!"+fmt.Sprint(leftColIndex)]
		if exists3 {
			results = append(results, val3)
		}
		val4, exists4 := numberMap[strconv.Itoa(rowAsInt)+"!"+fmt.Sprint(rightColIndex)]
		if exists4 {
			results = append(results, val4)
		}
		val5, exists5 := numberMap[strconv.Itoa(topRowIndex)+"!"+fmt.Sprint(rightColIndex)] 
		if exists5 {
			results = append(results, val5)
		}
		val6, exists6 := numberMap[strconv.Itoa(topRowIndex)+"!"+fmt.Sprint(leftColIndex)]
		if exists6 {
			results = append(results, val6)
		}
		val7, exists7 := numberMap[strconv.Itoa(bottomRowIndex)+"!"+fmt.Sprint(rightColIndex)]
		if exists7 {
			results = append(results, val7)
		}
		val8, exists8 := numberMap[strconv.Itoa(bottomRowIndex)+"!"+fmt.Sprint(leftColIndex)]
		if exists8{
			results = append(results, val8)
		}

		// watch me whip... watch me nae nae...
		removed := removeDuplicateInt(results)

		if len(removed) == 2 {sum = sum + (removed[0] * removed[1])}
	}

	return sum
}

func removeDuplicateInt(intSlice []int) []int {
    allKeys := make(map[int]bool)
    list := []int{}
    for _, item := range intSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func parseInput(input string) (symbolMap map[string]string, numberMap map[string]int) {
	numberMap = make(map[string]int)
	symbolMap = make(map[string]string)

	for lineIndex, line := range strings.Split(input, "\n") {
		re:=regexp.MustCompile(`\d+`)
		numberPos := re.FindAllStringIndex(line, -1)
		numbers := re.FindAllString(line, -1)

		for index, pos := range numberPos {
			posAsStr := strconv.Itoa(lineIndex) + "!" + strconv.Itoa(pos[0]) + "?" + strconv.Itoa(pos[1] - 1)
			numberMap[posAsStr] = stringToInt(numbers[index])
		}

		reS := regexp.MustCompile(`[^\w\s.]|\*`)
		symbolPos := reS.FindAllStringIndex(line, -1)
    	symbols := reS.FindAllString(line, -1)

		for index, sPos := range symbolPos {
			sPosAsStr := strconv.Itoa(lineIndex) + "!" + strconv.Itoa(sPos[0])
			symbolMap[sPosAsStr] = symbols[index]
		}
	}

	return symbolMap, numberMap
}

func parseInputPart2(input string) (symbolMap map[string]string, numberMap map[string]int) {
	numberMap = make(map[string]int)
	symbolMap = make(map[string]string)

	for lineIndex, line := range strings.Split(input, "\n") {
		re:=regexp.MustCompile(`\d+`)
		numberPos := re.FindAllStringIndex(line, -1)
		numbers := re.FindAllString(line, -1)

		for index, pos := range numberPos {
			for i := pos[0]; i <= pos[1] - 1; i++ {
			posAsStr := strconv.Itoa(lineIndex) + "!" + strconv.Itoa(i)
			numberMap[posAsStr] = stringToInt(numbers[index])
			}
		}

		reS := regexp.MustCompile(`\*`)
		symbolPos := reS.FindAllStringIndex(line, -1)
    	symbols := reS.FindAllString(line, -1)

		for index, sPos := range symbolPos {
			sPosAsStr := strconv.Itoa(lineIndex) + "!" + strconv.Itoa(sPos[0])
			symbolMap[sPosAsStr] = symbols[index]
		}
	}

	return symbolMap, numberMap
}


func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}

// too low 139373
// too low 525644
// too low 530903

// p2 
// too low 35207839