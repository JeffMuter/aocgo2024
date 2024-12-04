package daytwo

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

		// loop turns slice of strings into slice of ints
		for i := range splitLineString {
			convertedInt, err := strconv.Atoi(splitLineString[i])
			if err != nil {
				return -1, fmt.Errorf("error converting string to int: %w\n", err)
			} else if convertedInt == 0 {
				return -1, fmt.Errorf("error converting string to int: %w\ncurrent string val:'%s'", err, splitLineString[i])
			}
			splitLineInt = append(splitLineInt, convertedInt)
		}

		// assert that the split slice of its isnt 0, and also the same length as expected.
		if len(splitLineInt) != len(splitLineString) {
			return -1, fmt.Errorf("error converting string slice to int slice. unequal lengths between them\n")
		}

		err = checkForSafe(splitLineInt)
		if err != nil { // begin process to check if removing one number can be safe.
			err = checkForNearlySafe(splitLineInt)
			if err != nil {
				fmt.Printf("line completely unsafe: %v\nline: %v\n\n", err, line)
				continue
			}
			fmt.Printf("line is barely safe: %v\n\n", line)
			safetyScore++
			continue
		}
		fmt.Println("line completely safe: ", line)
		fmt.Println()
		safetyScore++
	}

	return safetyScore, nil
}

func checkForIncreasing(intSlice []int) error {
	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] < intSlice[i+1] {
			// conditional check to make sure difference not greater than 3
			if intSlice[i+1]-intSlice[i] > 3 {
				return fmt.Errorf("difference between too great while increasing...\n")
			}
			continue
		}
		return fmt.Errorf("was not incrementing... not safe...")
	}
	return nil
}

func checkForDecreasing(intSlice []int) error { // only called when index0 is confirmed less and i1
	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] > intSlice[i+1] {
			// conditional check to make sure difference not greater than 3
			if intSlice[i]-intSlice[i+1] > 3 {
				return fmt.Errorf("difference between too great while decreasing...\n")
			}
			continue
		}
		return fmt.Errorf("was not decrementing, not safe...\n")
	}
	return nil
}

func removeIndex(index int, slice []int) []int {
	newSlice := make([]int, len(slice)) // Create a copy of the slice
	copy(newSlice, slice)               // Copy the elements into the new slice
	return append(newSlice[:index], newSlice[index+1:]...)
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
	} else if slice[0] == slice[1] {
		return fmt.Errorf("error checking for safe: index 0 and 1 are tied.")
	}
	return nil
}

func checkForNearlySafe(slice []int) error {
	// loop through the slice, removing one element at a time, calling checkforsafe on each variation if any are safe, return nil, otherwise, if all variations arent safe, return err
	for i := 0; i < len(slice); i++ {
		var poppedSlice []int
		poppedSlice = removeIndex(i, slice)
		fmt.Printf("slice popped: %v... altered variant: %v\n", slice, poppedSlice)
		err := checkForSafe(poppedSlice)
		if err != nil {
			continue
		}
		fmt.Printf("altered line safe: %v... original: %v\n", poppedSlice, slice)
		return nil // this can only trigger if slice is safe, so return nil
	}
	return fmt.Errorf("no version of this slice was safe...\n")
}
