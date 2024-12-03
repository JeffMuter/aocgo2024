package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayTwoWork() (int, error) {
	fmt.Println("working on day 2...")

	var safetyScore int

	file, err := os.Open("days/dayTwoData.txt")
	if err != nil {
		fmt.Println("could not open file, big sad")
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// split each line into an array of 5 int.

		line := scanner.Text() // current line contents as string

		splitLineString := strings.Split(line, " ") // splits into slice of ["a", "b", "c", "d", "e"]string
		var splitLineInt []int

		for i := range splitLineString {
			convertedInt, err := strconv.Atoi(splitLineString[i])
			if err != nil {
				return -1, fmt.Errorf("error converting string to int: %w\n", err)
			} else if convertedInt == 0 {
				return -1, fmt.Errorf("error converting string to int: %w\ncurrent string val:'%s'", err, splitLineString[i])
			}
			splitLineInt = append(splitLineInt, convertedInt)
		}

		if len(splitLineInt) == 0 {
			return -1, fmt.Errorf("error converting strings to ints: %w\nline: %s/n", err, line)
		}

	}

	return safetyScore, nil
}

func checkForIncreasing(intSlice []int) error {
	var phoenix bool
	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] < intSlice[i+1] {
			// conditional check to make sure difference not greater than 3
			if intSlice[i+1]-intSlice[i] > 3 {
				fmt.Printf("revival status: index is: %d\n", i)
				if !phoenix {
					removeIndex(i, intSlice)
					phoenix = true
					i--
					continue
				}
				return fmt.Errorf("difference between too great while increasing...")
			}
			continue
		} else {
			if !phoenix {
				fmt.Printf("revival status: index is: %d\n", i)
				removeIndex(i, intSlice)
				phoenix = true
				i--
				continue
			}
			return fmt.Errorf("was not incrementing... not safe...")
		}
	}

	return nil
}

func checkForDecreasing(intSlice []int) error { // only called when index0 is confirmed less and i1
	var phoenix bool
	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] > intSlice[i+1] {
			// conditional check to make sure difference not greater than 3
			if intSlice[i]-intSlice[i+1] > 3 {
				if !phoenix {
					fmt.Printf("revival status: index is: %d\n", i)
					removeIndex(i, intSlice)
					phoenix = true
					i--
					continue
				}
				return fmt.Errorf("difference between too great while decreasing...")
			}
			continue
		} else {
			if !phoenix {
				fmt.Printf("revival status: index is: %d\n", i)
				removeIndex(i, intSlice)
				phoenix = true
				i--
				continue
			}
			return fmt.Errorf("was not decrementing... not safe...")
		}
	}

	return nil
}

func removeIndex(index int, slice []int) []int {
	fmt.Printf("removing index: [%d]\nslice\nline: %v\n", index, slice)
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("slice after removal: %v\n", slice)
	return slice
}

func checkForSafe(slice []int) error {

	if slice[0] < slice[1] {
		err := checkForIncreasing(slice)
		if err != nil {
			return fmt.Errorf("error:increasing value proved unsafe... slice used: %v\n err value: %w\n", slice, err)
		}
	} else if slice[0] > slice[1] {
		err := checkForDecreasing(slice)
		if err != nil {
			return fmt.Errorf("error:decreasing value proved unsafe... slice used: %v\n err value: %w\n", slice, err)
		}
	}
	return nil
}
