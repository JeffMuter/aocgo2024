package daythree

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateResult() (int, error) {

	var result int

	text, err := getChallengeStringFromFile()
	if err != nil {
		return -1, fmt.Errorf("could not open file to create result on day three: %w\n", err)
	}

	delimiter := "mul("

	splitByPrepend := strings.Split(text, delimiter)

	fmt.Printf("length of split: %d\nsplit value: %v\n\n", len(splitByPrepend), splitByPrepend)

	for i := 0; i < len(splitByPrepend); i++ {

		fmt.Println(splitByPrepend[i])
		if len(splitByPrepend[i]) < 4 {
			fmt.Printf("continue, assuming string is not long enough to contain valid input: %s\n", splitByPrepend[i])
			continue
		}
		before := splitByPrepend[i]

		afterRemMul := splitByPrepend[i]

		indexOfParenthesisEnd := strings.Index(splitByPrepend[i], ")")

		// if either end parenthesis missing, invalid, remove
		if indexOfParenthesisEnd == -1 {
			splitByPrepend = append(splitByPrepend[:i], splitByPrepend[i+1:]...)
			i--
			fmt.Println("couldnt find end parenth, ending for this string...")
			continue
		}

		// remove everything after ')'
		currentStringNoEndParenthesis := splitByPrepend[i][:indexOfParenthesisEnd]
		splitByPrepend[i] = currentStringNoEndParenthesis

		// check for existing comma here, grab its index for splitting? maybe better way to validate during splitting.
		indexOfComma := strings.Index(splitByPrepend[i], ",")
		if indexOfComma == -1 {
			splitByPrepend = append(splitByPrepend[:i], splitByPrepend[i+1:]...)
			i--
			fmt.Println("couldnt find index of comma continuing...")
			continue
		}

		splitPairForMultiplication := strings.Split(splitByPrepend[i], ",")
		if len(splitPairForMultiplication) != 2 {
			splitByPrepend = append(splitByPrepend[:i], splitByPrepend[i+1:]...)
			i--
			fmt.Println("length of split pair not 2... invalid... contuiing...")
			continue
		}

		firstInt, err := strconv.Atoi(splitPairForMultiplication[0])
		if err != nil {
			fmt.Println("first of multiplication pair was not an int..." + splitPairForMultiplication[0])
			splitByPrepend = append(splitByPrepend[:i], splitByPrepend[i+1:]...)
			i--
			continue
		}

		secondInt, err := strconv.Atoi(splitPairForMultiplication[1])
		if err != nil {
			fmt.Println("second of multiplication pair was not an int..." + splitPairForMultiplication[1])
			splitByPrepend = append(splitByPrepend[:i], splitByPrepend[i+1:]...)
			i--
			continue
		}

		result += firstInt * secondInt

		fmt.Printf(
			"initial string: %s\n"+
				"after removing mul(: %s\n"+
				"index of endPar: %d\n"+
				"commaIndex: %d\n"+
				"finalstringvalue: %s\n"+
				"firstSplitOfPair: %s\n"+
				"second: %s\n\n",

			before,
			afterRemMul,
			indexOfParenthesisEnd,
			indexOfComma,
			splitByPrepend[i],
			splitPairForMultiplication[0],
			splitPairForMultiplication[1],
		)
	}

	//	splitByPostend := strings.SplitAfter(strippedText, ")")

	//	for i := range splitByPostend {
	//		fmt.Printf("firstStrip: %s\nsecondSplit: %s\n\n", splitByPrepend[i], splitByPostend[i])
	//	}

	if result == 0 {
		return -1, fmt.Errorf("result is 0, check for input issues... %v", text)
	}

	return result, nil
}

func splitBefore(input, delimiter string) ([]string, error) {
	var result []string

	splitSlice := strings.SplitAfter(input, delimiter)
	fmt.Printf("splitbefore log:\nlength: %d\n0th: %s\n1st: %s\n", len(splitSlice), splitSlice[0], splitSlice[1])

	return result, nil
}

func getChallengeStringFromFile() (string, error) {

	// find file, open file, read file into string
	var fileLocation = "days/daythree/data.txt"
	file, err := os.Open(fileLocation)
	if err != nil {
		return "", fmt.Errorf("could not open file... %w\n@ loc: %s", err, fileLocation)
	}
	defer file.Close()

	text, err := os.ReadFile("days/daythree/data.txt")

	return string(text), nil
}

// split the whole file by 'mul(' ["mul(sdlfjksdfsfsf)", "mul(23,44)", "asdlfkj", "sfsdfs"]

// check if each substring contains a ) at any point.

// get what's in between mul(sdsdfdf, 234ljkh*) -->sdsdfdf, 234ljkh*

// check for comma, if a comma, yay. if not, err.

// split based on comma. make this the method of the last step if possible

// now we have 2 strings ["sdsdfdf", "234ljkh*"]

// check if they're both integers. if yes, good to multiply together.

// if not, err.
