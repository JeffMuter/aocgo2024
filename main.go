package main

import (
	"aocgo2024/days"
	"fmt"
)

func main() {
	result, err := days.DayTwoWork()
	if err != nil {
		fmt.Printf("error in getting locId from data: %v\n", err)
	}
	fmt.Println(result)
}
