package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	points_folds := strings.Split(string(file), "\n\n")

	max_x := 10000

	paper := make(map[int]int)

	for _, point := range strings.Split(points_folds[0], "\n") {

		sp := strings.Split(point, ",")

		x, _ := strconv.Atoi(sp[0])
		y, _ := strconv.Atoi(sp[1])

		val := max_x*x + y

		paper[val] = 1
	}

	for _, fold := range strings.Split(points_folds[1], "\n") {
		var axis byte
		var fold_value int

		fmt.Println(fold)
		fmt.Sscanf(fold, "fold along %c=%d", &axis, &fold_value)

		vals := []int{}

		for val, _ := range paper {

			vals = append(vals, val)

		}

		for _, val := range vals {

			x := val / max_x

			y := val % max_x

			if axis == 'x' {
				if x > fold_value {
					xd := x - fold_value
					x = fold_value - xd
					delete(paper, val)

					paper[x*max_x+y] = 1

				}
			} else {
				if y > fold_value {
					yd := y - fold_value
					y = fold_value - yd
					delete(paper, val)
					paper[x*max_x+y] = 1
				}
			}

		}

	}

	rect := image.Rect(-100, -100, 100, 100)
	img := image.NewRGBA(rect)

	green := color.RGBA{0, 255, 0, 255}
	draw.Draw(img, rect, &image.Uniform{green}, image.ZP, draw.Src)

	red := color.RGBA{255, 0, 0, 255}
	for val, _ := range paper {
		x := val / max_x
		y := val % max_x

		img.Set(x, y, red)

		fmt.Printf(" %d %d setting this point\n", x, y)
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)

}
