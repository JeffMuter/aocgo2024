package daythree

import (
	"fmt"
	"os"
	"strings"
)

func CreateResult() (int, error) {

	text, err := getChallengeStringFromFile()
	if err != nil {
		return -1, fmt.Errorf("could not open file to create result on day three: %w\n", err)
	}

	delimiter := "mul("
	var strippedText string

	splitByPrepend := strings.SplitAfter(text, delimiter)

	for i := range splitByPrepend {
		fmt.Println(splitByPrepend[i])
		strippedText += splitByPrepend[i]
	}

	splitByPostend := strings.SplitAfter(strippedText, ")")

	for i := range splitByPostend {
		fmt.Printf("line: %s	%s\n", splitByPrepend[i], splitByPostend[i])
	}

	return -1, nil
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
