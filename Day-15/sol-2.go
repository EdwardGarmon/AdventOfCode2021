package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	dirs := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	cost_grid := [][]int{}

	for _, row := range strings.Split(string(file), "\n") {

		cost_row := []int{}
		for _, col := range row {
			val, _ := strconv.Atoi(string(col))

			cost_row = append(cost_row, val)
		}

		cost_grid = append(cost_grid, cost_row)
	}

	dist_grid := [][]int{}

	max_y := len(cost_grid) - 1
	max_x := len(cost_grid) - 1

	get_cost := func(x int, y int) int {

		mx := x % (max_x + 1)
		my := y % (max_y + 1)

		// println(my, "x value modded")

		orig := cost_grid[mx][my]

		orig += x / (max_x + 1)
		orig += y / (max_y + 1)

		if orig > 9 {
			orig %= 9
		}

		return orig
	}

	front := make(MinHeap, 0)

	big_cost := [][]int{}

	for x := 0; x < (max_x+1)*5; x++ {
		big_row := []int{}
		dist_row := []int{}

		for y := 0; y < (max_y+1)*5; y++ {

			big_row = append(big_row, get_cost(x, y))
			dist_row = append(dist_row, math.MaxInt32)

		}
		big_cost = append(big_cost, big_row)
		dist_grid = append(dist_grid, dist_row)
	}

	big_cost[0][0] = 1

	get_tiles := func(x int, y int) [][]int {

		tiles := [][]int{}
		for _, dir := range dirs {
			nx := x + dir[0]
			ny := y + dir[1]

			if nx < 0 || ny < 0 || nx >= len(big_cost) || ny >= len(big_cost[0]) {
				continue
			}

			// println(x, y, "in get tiles")
			tiles = append(tiles, []int{nx, ny})

		}
		return tiles
	}

	heap.Push(&front, &Item{x: 0, y: 0, value: 0})

	for true {
		// select a new node
		pop := heap.Pop(&front)
		lowest := pop.(*Item)

		if dist_grid[lowest.x][lowest.y] < lowest.value {
			continue
		}

		if lowest.x == len(big_cost)-1 && lowest.y == len(big_cost)-1 {
			println(dist_grid[lowest.x][lowest.y])
			break
		}

		// for all of the neigbors of this node
		for _, neigh := range get_tiles(lowest.x, lowest.y) {
			// println(neigh[0], neigh[1], "in loop")
			next := Item{
				value: big_cost[neigh[0]][neigh[1]] + lowest.value,
				x:     neigh[0],
				y:     neigh[1],
			}

			if next.value < dist_grid[next.x][next.y] {
				heap.Push(&front, &next)
				dist_grid[next.x][next.y] = next.value

			}

		}
	}

	// for _, row := range big_cost {
	// 	for _, val := range row {
	// 		print(val)
	// 	}
	// 	println()
	// }

}

type Item struct {
	value int
	x     int
	y     int
}

type MinHeap []*Item

func (mh MinHeap) Len() int { return len(mh) }

func (pq *MinHeap) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq MinHeap) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
