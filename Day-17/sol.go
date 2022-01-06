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

	ranges := string(file)

	println(ranges)

	var min_x, min_y, max_x, max_y int

	fmt.Sscanf(ranges, "target area: x=%d..%d, y=%d..%d",
		&min_x, &max_x, &min_y, &max_y)

	println("max y value: ", ((min_y+1)*min_y)/2)

	x_t_vals := make(map[int][]int)

	y_t_vals := make(map[int][]int)

	max_time := 0

	for i := min_y; i <= -(min_y + 1); i++ {
		pos := 0
		time := 0

		for y := i; pos > min_y; y-- {
			pos += y
			if pos >= min_y && pos <= max_y {
				// println(i, time)
				addToArrMap(i, time, y_t_vals)
			}

			time++
		}
		if time > max_time {
			max_time = time
		}

	}

	for i := 0; i <= max_x; i++ {

		pos := 0
		time := 0

		x := i

		for time <= max_time {

			pos += x
			if pos <= max_x && pos >= min_x {
				addToArrMap(i, time, x_t_vals)
			}
			if x > 0 {
				x--
			}
			time++

		}
	}

	combos := 0

	sol := make(map[string]int)

	for y_time, y_array := range y_t_vals {
		for x_time, x_array := range x_t_vals {

			if y_time == x_time {
				for _, xv := range x_array {
					for _, yv := range y_array {
						combos++
						s := fmt.Sprintf("%d,%d", xv, yv)
						sol[s] = 0
					}
				}
			}

		}

	}

	println(len(sol))

}

func incrementMap(ch int, count_map map[int]int) {

	count, exists := count_map[ch]

	if exists {
		count_map[ch] = count + 1
	} else {
		count_map[ch] = 1
	}

}

func addToArrMap(i int, time int, x_t_vals map[int][]int) {

	t_arr, exists := x_t_vals[time]

	if exists {
		x_t_vals[time] = append(t_arr, i)

	} else {
		x_t_vals[time] = []int{i}
	}
}
