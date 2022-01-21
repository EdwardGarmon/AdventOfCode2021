package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}
type Scanner struct {
	points []Point
}

type Dist_Mat []map[float64]int

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	scanners := []Scanner{}

	for _, scanner := range strings.Split(string(file), "\n\n") {

		scan := Scanner{}
		scan.points = []Point{}

		for i, point := range strings.Split(scanner, "\n") {
			if i == 0 {
				continue
			}

			xyz := strings.Split(point, ",")
			npoint := Point{}

			x, _ := strconv.Atoi(xyz[0])
			y, _ := strconv.Atoi(xyz[1])
			z, _ := strconv.Atoi(xyz[2])

			npoint.x = x
			npoint.y = y
			npoint.z = z

			scan.points = append(scan.points, npoint)
		}

		scanners = append(scanners, scan)
	}

	visited := []bool{}

	for _, _ = range scanners {
		visited = append(visited, false)
	}

	num := 1

	mainScanner := scanners[0]
	visited[0] = true

	index := 1

	scaner_locs := []Point{{0, 0, 0}}

	for num < len(scanners) {

		for visited[index] {
			index++
			if index >= len(scanners) {
				index = 1
			}
		}

		newSc, foundNew, scan_loc := mainScanner.mergeScanner(scanners[index])

		if foundNew {
			scaner_locs = append(scaner_locs, scan_loc)
			mainScanner = newSc
			visited[index] = true
			num++

		}

		index++
		if index >= len(scanners) {
			index = 1
		}

	}

	println(len(mainScanner.points))

	max := 0

	for i := 0; i < len(scaner_locs); i++ {
		for x := i; x < len(scaner_locs); x++ {
			d := scaner_locs[x].manhattan(scaner_locs[i])

			if d > max {
				max = d
			}
		}
	}

	println("max dist", max)

}

func (scan1 Scanner) mergeScanner(scan2 Scanner) (Scanner, bool, Point) {

	m1 := scan1.distance_mat()
	m2 := scan2.distance_mat()

	points := common_points(m1, m2)

	orient, xt, yt, zt := findOrientTrans(points, scan1, scan2)

	newPoints := []Point{}

	for _, p := range scan1.points {
		newPoints = append(newPoints, p)
	}

	if orient == -1 {
		return Scanner{newPoints}, false, Point{}
	}

	for index, point := range scan2.points {

		needToAdd := true

		for _, pindex := range points {
			if pindex[1] == index {
				needToAdd = false
			}
		}

		if needToAdd {
			newPoint := point.orientations()[orient]
			newPoint.x += xt
			newPoint.y += yt
			newPoint.z += zt
			newPoints = append(newPoints, newPoint)
		}

	}

	return Scanner{newPoints}, true, Point{xt, yt, zt}
}

func findOrientTrans(points [][]int, scanner1 Scanner, scanner2 Scanner) (int, int, int, int) {
	if len(points) == 0 {
		return -1, 0, 0, 0
	}
	for i := 0; i < 24; i++ {

		pair := points[0]

		p1 := scanner1.points[pair[0]]
		p2 := scanner2.points[pair[1]].orientations()[i]

		xt := p1.x - p2.x
		yt := p1.y - p2.y
		zt := p1.z - p2.z

		for n := 1; n < 4; n++ {

			pair := points[n]
			p1 := scanner1.points[pair[0]]
			p2 := scanner2.points[pair[1]].orientations()[i]

			nx := p1.x - p2.x
			ny := p1.y - p2.y
			nz := p1.z - p2.z

			if nx != xt || ny != yt || zt != nz {
				break
			}

			if n == 2 {
				return i, xt, yt, zt
			}

		}

	}

	return -1, 0, 0, 0
}

func common_points(m1 Dist_Mat, m2 Dist_Mat) [][]int {
	for _, distMap1 := range m1 {
		for _, distMap2 := range m2 {

			matches := [][]int{}

			for dist, index := range distMap1 {

				value, exists := distMap2[dist]

				if exists {
					matches = append(matches, []int{index, value})
				}
			}

			if len(matches) >= 12 {
				return matches
			}

		}
	}
	return [][]int{}
}

func (sc Scanner) distance_mat() Dist_Mat {

	mat := []map[float64]int{}

	for _, p1 := range sc.points {
		m := make(map[float64]int)
		for n2, p2 := range sc.points {

			m[p1.dist(p2)] = n2
		}
		mat = append(mat, m)
	}

	return mat
}

func (p1 Point) dist(p2 Point) float64 {
	return math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2)
}

func (p1 Point) manhattan(p2 Point) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x))) +
		int(math.Abs(float64(p1.y)-float64(p2.y))) +
		int(math.Abs(float64(p1.z)-float64(p2.z)))
}

//this is probably the least efficient part
//need to precompute these
func (p Point) orientations() []Point {

	parr := []int{p.x, p.y, p.z}

	orients := []Point{}

	axis := []int{0, 1, 2}
	dir := []int{-1, 1}

	for _, i := range axis {
		rem1 := []int{}

		for _, ox := range axis {
			if ox != i {
				rem1 = append(rem1, ox)
			}
		}

		for _, dx := range dir {

			for _, j := range rem1 {
				for _, dy := range dir {

					rem := []int{}

					for _, ox := range rem1 {
						if ox != j {
							rem = append(rem, ox)
						}
					}

					z := parr[rem[0]] * cross_sign(i, j, dx, dy)

					orients = append(orients, Point{parr[i] * dx, parr[j] * dy, z})

				}
			}
		}
	}
	return orients

}

func cross_sign(i int, j int, dx int, dy int) int {

	if i == 0 {
		if j == 1 {
			return dx * dy
		}

		if j == 2 {
			return dx * dy * -1
		}
	}

	if i == 1 {
		if j == 0 {
			return dx * dy * -1
		}

		if j == 2 {
			return dx * dy
		}
	}

	if i == 2 {

		if j == 0 {
			return dx * dy
		}

		if j == 1 {
			return dx * dy * -1
		}

	}
	return 10000
}
