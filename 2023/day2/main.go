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

	fmt.Printf("The sum of the IDs of the possible games is: %d\n", possibleGameIDSum)
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
