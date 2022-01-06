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

	last := -1

	incs := 0

	for _, d := range strings.Split(string(file), "\n") {

		cur, err := strconv.Atoi(d)

		if err != nil {
			return
		}

		if last == -1 {
			last = cur
			continue
		}

		if cur > last {
			incs++
		}

		last = cur

	}

	fmt.Println(incs)
}
