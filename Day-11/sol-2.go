package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	var grid [10][10]int

	flashed := make(map[int]int)

	for row, d := range strings.Split(string(file), "\n") {

		for col, value := range strings.Split(d, "") {

			num, _ := strconv.Atoi(value)

			grid[row][col] = num

		}

	}

	flashes := 0

	var check_for_flash func(int, int)

	check_for_flash = func(row int, cin int) {

		index := row*10 + cin

		_, has_flashed := flashed[index]

		if grid[row][cin] > 9 && !has_flashed {

			// fmt.Printf("flashed %d %d %d \n", row, cin, grid[row][cin])

			flashed[index] = 1

			flashes++

			for _, x := range []int{-1, 0, 1} {
				for _, y := range []int{-1, 0, 1} {

					if x == 0 && y == 0 {
						continue
					}

					xn := row + x
					yn := cin + y

					if xn < 0 || xn > 9 || yn < 0 || yn > 9 {
						continue
					}

					grid[xn][yn]++
					check_for_flash(xn, yn)

				}
			}
		}
	}

	//increase energy level by one

	simul := false
	steps := 0

	for !simul {
		for row, col := range grid {
			for cin := range col {
				grid[row][cin]++

			}
		}

		for row, col := range grid {
			for cin := range col {

				index := row*10 + cin

				_, has_flashed := flashed[index]

				if !has_flashed {
					check_for_flash(row, cin)
				}

			}
		}

		if len(flashed) == 100 {
			simul = true
		}

		for key := range flashed {
			row := key / 10
			cin := key % 10

			// fmt.Printf("retrieved flashed %d %d %d \n", row, cin, key)
			grid[row][cin] = 0

		}

		for row, col := range grid {
			for cin := range col {

				fmt.Print(grid[row][cin])

			}
			fmt.Println()
		}
		fmt.Println()

		// fmt.Println("end of iteration")
		flashed = make(map[int]int)
		steps++
	}

	fmt.Println(steps)

}
