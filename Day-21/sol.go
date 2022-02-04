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

	format := "Player 1 starting position: %d\nPlayer 2 starting position: %d"

	position := []int{0, 0}

	fmt.Sscanf(string(file), format, &position[0], &position[1])

	println(position[0], position[1])

	dice := 1
	rolls := 0
	index := 0

	scores := []int{0, 0}

	for scores[0] < 1000 && scores[1] < 1000 {

		// println(dice, "old")
		move, n_dice := roll_dice(dice)
		dice = n_dice
		// println(n_dice, "new")

		old := position[index]

		position[index] += (move % 10)

		if position[index] > 10 {

			position[index] %= 10

		}

		scores[index] += position[index]

		// if rolls > 1323 {

		println("player ", index, " moved ", move, " from space ", old, "to space ", position[index], "score ", scores[index])

		rolls += 3

		index++
		index %= 2

	}

	value := scores[0]

	if scores[0] > scores[1] {
		value = scores[1]
	}

	println(rolls * value)

}

func roll_dice(current int) (int, int) {

	total := 0
	for i := 0; i < 3; i++ {
		total += current
		current += 1
		if current > 100 {
			current = 1
		}
	}

	return total, current

}
