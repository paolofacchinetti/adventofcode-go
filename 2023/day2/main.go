package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/paolofacchinetti/adventofcode-go/utils"
)

/*
solving this one without regex because they are boring
*/

func main() {
	/*
		Part 1
	*/
	file, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	gameNumber := 0
	possibleGameIDSum := 0

	for scanner.Scan() {
		gameNumber++
		line := scanner.Text()

		if IsGamePossible(line, 12, 13, 14) {
			possibleGameIDSum += gameNumber
		}

	}

	fmt.Printf("Part 1: the sum of the IDs of the possible games is: %d\n", possibleGameIDSum)

	/*
		Part 2
	*/
	file, err = utils.ReadInput()
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)
	sumOfPowers := 0

	for scanner.Scan() {
		sumOfPowers += GetPowerOfMinimumCubes(scanner.Text())
	}

	fmt.Printf("Part 2: the sum of powers of the minimum number of cubes per game is: %d\n", sumOfPowers)
}

func IsGamePossible(line string, maxRed int, maxGreen int, maxBlue int) bool {
	// remove "game N: "
	line = line[strings.Index(line, ":")+2:]
	// split game into sets
	sets := (strings.Split(line, ";"))
	for _, set := range sets {
		set = strings.TrimSpace(set)
		// split set into draws
		draws := strings.Split(set, ",")
		for _, draw := range draws {
			draw = strings.TrimSpace(draw)
			splitDraw := strings.Split(draw, " ")
			strNumber, color := splitDraw[0], splitDraw[1]
			number, _ := strconv.Atoi(strNumber)

			// check if number > max
			switch color {
			case "green":
				if number > maxGreen {
					return false
				}
			case "red":
				if number > maxRed {
					return false
				}
			case "blue":
				if number > maxBlue {
					return false
				}
			}
		}
	}
	return true
}

func GetPowerOfMinimumCubes(line string) int {
	minRed := 1
	minGreen := 1
	minBlue := 1
	// remove "game N: "
	line = line[strings.Index(line, ":")+2:]
	// split game into sets
	sets := (strings.Split(line, ";"))
	for _, set := range sets {
		// split set into draws
		draws := strings.Split(set, ",")
		for _, draw := range draws {
			draw = strings.TrimSpace(draw)
			splitDraw := strings.Split(draw, " ")
			strNumber, color := splitDraw[0], splitDraw[1]
			number, _ := strconv.Atoi(strNumber)

			// check if number > max
			switch color {
			case "green":
				if number > minGreen {
					minGreen = number
				}
			case "red":
				if number > minRed {
					minRed = number
				}
			case "blue":
				if number > minBlue {
					minBlue = number
				}
			}
		}
	}
	return minRed * minGreen * minBlue
}
