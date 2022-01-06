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

	dist_grid := [][]int{}
	cost_grid := [][]int{}

	for _, row := range strings.Split(string(file), "\n") {
		dist_row := []int{}
		cost_row := []int{}
		for _, col := range row {
			val, _ := strconv.Atoi(string(col))
			dist_row = append(dist_row, math.MaxInt32)
			cost_row = append(cost_row, val)
		}
		dist_grid = append(dist_grid, dist_row)
		cost_grid = append(cost_grid, cost_row)
	}

	max_y := len(cost_grid) - 1
	max_x := len(dist_grid) - 1

	println(max_x, max_y)

	get_tiles := func(x int, y int) [][]int {

		tiles := [][]int{}
		for _, dir := range dirs {
			nx := x + dir[0]
			ny := y + dir[1]

			if nx > max_x || ny > max_y ||
				nx < 0 || ny < 0 {
				continue
			}

			tiles = append(tiles, []int{nx, ny})

		}
		return tiles
	}

	front := make(MinHeap, 0)

	cost_grid[0][0] = 0

	heap.Push(&front, &Item{x: 0, y: 0, value: 0})

	for true {
		// select a new node
		pop := heap.Pop(&front)
		lowest := pop.(*Item)

		if dist_grid[lowest.x][lowest.y] < lowest.value {
			continue
		}

		if lowest.x == max_x && lowest.y == max_y {
			println(dist_grid[lowest.x][lowest.y])
			break
		}

		// for all of the neigbors of this node
		for _, neigh := range get_tiles(lowest.x, lowest.y) {

			next := Item{
				value: cost_grid[neigh[0]][neigh[1]] + lowest.value,
				x:     neigh[0],
				y:     neigh[1],
			}

			if next.value < dist_grid[next.x][next.y] {
				heap.Push(&front, &next)
				dist_grid[next.x][next.y] = next.value

			}

		}
	}

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
