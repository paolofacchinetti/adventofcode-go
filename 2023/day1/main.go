package main

import (
	"bufio"
	"fmt"
	"unicode"

	"github.com/paolofacchinetti/adventofcode-go/utils"
)

func main() {
	file, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	// iterate over text lines
	for scanner.Scan() {
		line := scanner.Text()
		lastDigit := 0
		firstDigitFound := false
		// iterate over characters in a line
		for _, char := range line {
			if unicode.IsDigit(char) {
				if !firstDigitFound {
					sum += 10 * int(char-'0')
					firstDigitFound = true
				}
				lastDigit = int(char - '0')
			}
		}
		sum += lastDigit
		fmt.Println(sum)
	}
	fmt.Printf("the total sum of all calibration values is %d\n", sum)
}
