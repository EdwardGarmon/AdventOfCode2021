package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("1-Day.input.txt")

	if err != nil {
		return
	}

	total := 0

	incs := 0

	values := strings.Split(string(file), "\n")

	for i, d := range strings.Split(string(file), "\n") {

		cur, _ := strconv.Atoi(d)

		if i < 3 {
			total += cur
			continue
		}

		past, _ := strconv.Atoi(values[i-3])
		newTotal := total + cur - past

		if newTotal > total {
			incs++
		}

		total = newTotal

	}

	fmt.Println(incs)
}
