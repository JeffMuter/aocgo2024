package dayone

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FirstPartDayOne() (int, error) {
	var leftInts, rightInts []int
	var result int

	// read data file
	file, err := os.Open("aocdata/01day.txt")
	if err != nil {
		fmt.Println("could not open file, big sad")
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()                  // current line contents as string
		splitLine := strings.Split(line, "   ") // splits into slice of [x, y]int

		var leftInt, rightInt int

		leftInt, err = strconv.Atoi(splitLine[0])
		if err != nil {
			fmt.Println("failed to convert first element of split slice into int")
			fmt.Printf(":%s:\n", splitLine[0])
			return -1, err
		}

		rightInt, err = strconv.Atoi(splitLine[1])
		if err != nil {
			fmt.Println("failed to convert second element of split slice into int")
			return -1, err
		}

		leftInts = append(leftInts, leftInt)
		rightInts = append(rightInts, rightInt)
	}

	if len(leftInts) != len(rightInts) {
		fmt.Println("err: slices from left and right ints from file not equal.")
		return -1, err
	}

	// compare in a loop by index, subtracting into positive int

	sort.Ints(leftInts)
	sort.Ints(rightInts)

	for i := 0; i < len(leftInts); i++ {
		thisLineDifference := leftInts[i] - rightInts[i]
		if thisLineDifference < 0 {
			thisLineDifference = thisLineDifference * -1
		}
		fmt.Printf("line[%d]: left: %d right: %d difference: %d\n", i, leftInts[i], rightInts[i], thisLineDifference)
		result += thisLineDifference
	}
	return result, nil
}

func SecondPartDayOne() (int, error) {
	var leftInts, rightInts []int

	// read data file
	file, err := os.Open("days/firstDayData.txt")
	if err != nil {
		fmt.Println("could not open file, big sad")
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()                  // current line contents as string
		splitLine := strings.Split(line, "   ") // splits into slice of [x, y]int

		var leftInt, rightInt int

		leftInt, err = strconv.Atoi(splitLine[0])
		if err != nil {
			fmt.Println("failed to convert first element of split slice into int")
			fmt.Printf(":%s:\n", splitLine[0])
			return -1, err
		}

		rightInt, err = strconv.Atoi(splitLine[1])
		if err != nil {
			fmt.Println("failed to convert second element of split slice into int")
			return -1, err
		}

		leftInts = append(leftInts, leftInt)
		rightInts = append(rightInts, rightInt)
	}

	if len(leftInts) != len(rightInts) {
		fmt.Println("err: slices from left and right ints from file not equal.")
		return -1, err
	}

	var similarityScore int

	for i := 0; i < len(leftInts); i++ { // nested loop. for each item of left, loop through right for sim score
		var currentSimilarityScore int
		for j := 0; j < len(leftInts); j++ {
			if leftInts[i] == rightInts[j] {
				currentSimilarityScore++
			}
		}
		similarityScore += currentSimilarityScore * leftInts[i]
	}
	return similarityScore, nil
}
