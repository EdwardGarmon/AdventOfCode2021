package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	input := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

}

func validate_input(num []int) {

}
