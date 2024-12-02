package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/paolofacchinetti/adventofcode-go/utils"
)

func main() {
	file, err := utils.ReadInputAsFile()
	if err != nil {
		panic(err)
	}

	calcDistance(file)

}

func calcDistance(file *os.File) {

	scanner := bufio.NewScanner(file)
	var a, b []int
	dist := 0

	// iterate over text lines
	for scanner.Scan() {
		line := scanner.Text()
		// Split the string by spaces
		parts := strings.Fields(line)

		// Convert strings to integers
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		a = append(a, num1)
		b = append(b, num2)
	}

	slices.Sort(a)
	slices.Sort(b)

	for i := range a {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = -diff
		}
		dist += diff
	}

	fmt.Println(dist)
}
