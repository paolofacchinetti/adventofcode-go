package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/paolofacchinetti/adventofcode-go/utils"
)

func main() {
	file, err := utils.ReadInput()
	if err != nil {
		panic(err)
	}

	part1Calc(file)

	file, err = utils.ReadInput()
	if err != nil {
		panic(err)
	}

	part2Calc(file)
}

func part1Calc(file *os.File) {
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
	}
	fmt.Printf("Part 1: the total sum of all calibration values is %d\n", sum)
}

func part2Calc(file *os.File) {
	scanner := bufio.NewScanner(file)
	sum := 0

	// iterate over text lines
	for scanner.Scan() {
		line := scanner.Text()

		// map[index]number
		m := make(map[int]int)

		// index spelled out numbers
		indexSpelledOutNums(line, m)

		// index digits
		indexDigits(line, m)

		// get first digit
		for i := 0; i < len(line); i++ {
			if num, ok := m[i]; ok {
				sum += num * 10
				break
			}
		}

		// get last digit
		for i := len(line) - 1; i >= 0; i-- {
			if num, ok := m[i]; ok {
				sum += num
				break
			}
		}
	}
	fmt.Printf("Part 2: the total sum of all calibration values is %d\n", sum)
}

func indexDigits(line string, m map[int]int) {
	for pos, char := range line {
		if unicode.IsDigit(char) {
			m[pos] = int(char - '0')
		}
	}
}

func indexSpelledOutNums(line string, m map[int]int) {
	if i := strings.Index(line, "one"); i >= 0 {
		m[i] = 1
	}
	if i := strings.LastIndex(line, "one"); i >= 0 {
		m[i] = 1
	}
	if i := strings.Index(line, "two"); i >= 0 {
		m[i] = 2
	}
	if i := strings.LastIndex(line, "two"); i >= 0 {
		m[i] = 2
	}
	if i := strings.Index(line, "three"); i >= 0 {
		m[i] = 3
	}
	if i := strings.LastIndex(line, "three"); i >= 0 {
		m[i] = 3
	}
	if i := strings.Index(line, "four"); i >= 0 {
		m[i] = 4
	}
	if i := strings.LastIndex(line, "four"); i >= 0 {
		m[i] = 4
	}
	if i := strings.Index(line, "five"); i >= 0 {
		m[i] = 5
	}
	if i := strings.LastIndex(line, "five"); i >= 0 {
		m[i] = 5
	}
	if i := strings.Index(line, "six"); i >= 0 {
		m[i] = 6
	}
	if i := strings.LastIndex(line, "six"); i >= 0 {
		m[i] = 6
	}
	if i := strings.Index(line, "seven"); i >= 0 {
		m[i] = 7
	}
	if i := strings.LastIndex(line, "seven"); i >= 0 {
		m[i] = 7
	}
	if i := strings.Index(line, "eight"); i >= 0 {
		m[i] = 8
	}
	if i := strings.LastIndex(line, "eight"); i >= 0 {
		m[i] = 8
	}
	if i := strings.Index(line, "nine"); i >= 0 {
		m[i] = 9
	}
	if i := strings.LastIndex(line, "nine"); i >= 0 {
		m[i] = 9
	}
}
