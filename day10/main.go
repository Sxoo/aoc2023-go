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

type Pipe struct {
	coords string
	distanceFromStart int
	neighbours []string
	visited bool
	isInMainLoop bool
	isEnclosed bool
}

func part1(input string) int {
	parsed := parseInput(input)

	pipeMap := make(map[string]Pipe)

	var startCoords string

	for x, line := range parsed {
		for y, char := range line {
			if char == '.' {
				continue
			}

			if (string(char) == "S" ) {
				var neighbours []string
				startCoords = fmt.Sprint(x)+","+fmt.Sprint(y)

				// right -> '-' , '7' , 'J'
				// left -> '-' , 'L' , 'F'
				// top -> '|' , '7' , 'F'
				// bottom -> '|' , 'J' , 'L'

				if (y+1 < len(parsed[x])) {
					if (string(parsed[x][y+1]) == "-" || string(parsed[x][y+1]) == "7" || string(parsed[x][y+1]) == "J") {
						neighbours = append(neighbours, fmt.Sprint(x)+","+fmt.Sprint(y+1))
					}
				}
				if (y-1 >= 0) {
					if (string(parsed[x][y-1]) == "-" || string(parsed[x][y-1]) == "L" || string(parsed[x][y-1]) == "F") {
						neighbours = append(neighbours, fmt.Sprint(x)+","+fmt.Sprint(y-1))
					}
				}
				if (x+1 < len(parsed)) {
					if (string(parsed[x+1][y]) == "|" || string(parsed[x+1][y]) == "7" || string(parsed[x+1][y]) == "F") {
						neighbours = append(neighbours, fmt.Sprint(x+1)+","+fmt.Sprint(y))
					}
				}
				if (x-1 >= 0) {
					if (string(parsed[x-1][y]) == "|" || string(parsed[x-1][y]) == "J" || string(parsed[x-1][y]) == "L") {
						neighbours = append(neighbours, fmt.Sprint(x-1)+","+fmt.Sprint(y))
					}
				}

				pipeMap[fmt.Sprint(x)+","+fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), 0, neighbours, false, true, false}
				continue
			}

			if (string(char) == "|" ) {
				firstNeighborCoords := fmt.Sprint(x+1) + "," + fmt.Sprint(y)
				SecondNeighborCoords := fmt.Sprint(x-1) + "," + fmt.Sprint(y)

				pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},false, false, false}
				continue
			}

			if (string(char) == "-" ) {
				firstNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y+1)
				SecondNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y-1)

				pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},false, false, false}
				continue
			}

			if (string(char) == "L" ) {
				firstNeighborCoords := fmt.Sprint(x-1) + "," + fmt.Sprint(y)
				SecondNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y+1)

				pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},false, false, false}
				continue
			}

			if (string(char) == "J" ) {
				firstNeighborCoords := fmt.Sprint(x-1) + "," + fmt.Sprint(y)
				SecondNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y-1)

				pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords}, false,false, false}
				continue
			}

			if (string(char) == "7" ) {
				firstNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y-1)
				SecondNeighborCoords := fmt.Sprint(x+1) + "," + fmt.Sprint(y)

				pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},false,false, false}
				continue
			}

			if (string(char) == "F" ) {
				firstNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y+1)
				SecondNeighborCoords := fmt.Sprint(x+1) + "," + fmt.Sprint(y)

				pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},false,false, false}
				continue
			}
		}
	}



	coords := pipeMap[startCoords].neighbours[0]
	prevCoords := startCoords
	currDistance := 1

	maxDistance := 0

	for coords != startCoords {
		pipe := pipeMap[coords]
		pipe.visited = true
		pipe.distanceFromStart = currDistance

		for _, neighbour := range pipe.neighbours {
			if neighbour != coords && neighbour != prevCoords && !pipeMap[neighbour].visited{
				prevCoords = coords
				coords = neighbour
				break;
			}
		}

		if (currDistance > maxDistance) {
			maxDistance = currDistance
		}
	
		currDistance++
	}

	return (maxDistance + 1) / 2
}

func part2(input string) int {
parsed := parseInput(input)
fmt.Println(parsed)

	// pipeMap := make(map[string]Pipe)

	// var startCoords string

	// for x, line := range parsed {
	// 	for y, char := range line {
	// 		if char == '.' {
	// 			continue
	// 		}

	// 		if (string(char) == "S" ) {
	// 			var neighbours []string
	// 			startCoords = fmt.Sprint(x)+","+fmt.Sprint(y)

	// 			// right -> '-' , '7' , 'J'
	// 			// left -> '-' , 'L' , 'F'
	// 			// top -> '|' , '7' , 'F'
	// 			// bottom -> '|' , 'J' , 'L'

	// 			if (y+1 < len(parsed[x])) {
	// 				if (string(parsed[x][y+1]) == "-" || string(parsed[x][y+1]) == "7" || string(parsed[x][y+1]) == "J") {
	// 					neighbours = append(neighbours, fmt.Sprint(x)+","+fmt.Sprint(y+1))
	// 				}
	// 			}
	// 			if (y-1 >= 0) {
	// 				if (string(parsed[x][y-1]) == "-" || string(parsed[x][y-1]) == "L" || string(parsed[x][y-1]) == "F") {
	// 					neighbours = append(neighbours, fmt.Sprint(x)+","+fmt.Sprint(y-1))
	// 				}
	// 			}
	// 			if (x+1 < len(parsed)) {
	// 				if (string(parsed[x+1][y]) == "|" || string(parsed[x+1][y]) == "7" || string(parsed[x+1][y]) == "F") {
	// 					neighbours = append(neighbours, fmt.Sprint(x+1)+","+fmt.Sprint(y))
	// 				}
	// 			}
	// 			if (x-1 >= 0) {
	// 				if (string(parsed[x-1][y]) == "|" || string(parsed[x-1][y]) == "J" || string(parsed[x-1][y]) == "L") {
	// 					neighbours = append(neighbours, fmt.Sprint(x-1)+","+fmt.Sprint(y))
	// 				}
	// 			}

	// 			pipeMap[fmt.Sprint(x)+","+fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), 0, neighbours, false, true, false}
	// 			continue
	// 		}

	// 		if (string(char) == "|" ) {
	// 			firstNeighborCoords := fmt.Sprint(x+1) + "," + fmt.Sprint(y)
	// 			SecondNeighborCoords := fmt.Sprint(x-1) + "," + fmt.Sprint(y)

	// 			pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords}, false, false, false}
	// 			continue
	// 		}

	// 		if (string(char) == "-" ) {
	// 			firstNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y+1)
	// 			SecondNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y-1)

	// 			pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},  false,false, false}
	// 			continue
	// 		}

	// 		if (string(char) == "L" ) {
	// 			firstNeighborCoords := fmt.Sprint(x-1) + "," + fmt.Sprint(y)
	// 			SecondNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y+1)

	// 			pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},  false,false, false}
	// 			continue
	// 		}

	// 		if (string(char) == "J" ) {
	// 			firstNeighborCoords := fmt.Sprint(x-1) + "," + fmt.Sprint(y)
	// 			SecondNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y-1)

	// 			pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords},  false,false, false}
	// 			continue
	// 		}

	// 		if (string(char) == "7" ) {
	// 			firstNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y-1)
	// 			SecondNeighborCoords := fmt.Sprint(x+1) + "," + fmt.Sprint(y)

	// 			pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords}, false,false, false}
	// 			continue
	// 		}

	// 		if (string(char) == "F" ) {
	// 			firstNeighborCoords := fmt.Sprint(x) + "," + fmt.Sprint(y+1)
	// 			SecondNeighborCoords := fmt.Sprint(x+1) + "," + fmt.Sprint(y)

	// 			pipeMap[fmt.Sprint(x) + "," + fmt.Sprint(y)] = Pipe{fmt.Sprint(x)+","+fmt.Sprint(y), -1, []string{firstNeighborCoords, SecondNeighborCoords}, false,false, false}
	// 			continue
	// 		}
	// 	}
	// }



	// coords := pipeMap[startCoords].neighbours[0]
	// prevCoords := startCoords
	// currDistance := 1
	// maxDistance := 0

	// mainLoopMap := make(map[string]Pipe)

	// for coords != startCoords {
	// 	pipe := pipeMap[coords]
	// 	mainLoopMap[coords] = Pipe{coords, currDistance, pipe.neighbours, false, true, false}

	// 	for _, neighbour := range pipe.neighbours {
	// 		if neighbour != coords && neighbour != prevCoords{
	// 			prevCoords = coords
	// 			coords = neighbour
	// 			break;
	// 		}
	// 	}
	// }

	// // fmt.Println(mainLoopMap)

	// for x:= 0; x < 10; x++ {
	// 	for y:=0; y <20; y++ {
	// 		_, exists := mainLoopMap[fmt.Sprint(x)+","+fmt.Sprint(y)]
	// 		if exists {
	// 			fmt.Print("S")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
			
	// 	}
	// 	fmt.Print("\n")
	// }

	return 0
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