package main

import (
	"aocgo2024/daythree"
	"fmt"
)

func main() {
	result, err := daythree.CreateResult()
	if err != nil {
		fmt.Printf("error in getting locId from data: %v\n", err)
	}
	fmt.Println(result)
}
