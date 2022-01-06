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

	nums := strings.Split(string(file), "\n")

	max := 0

	for _, num := range nums {

		for _, num2 := range nums {
			num := parse(nil, num).(*snail_number)
			num2 := parse(nil, num2).(*snail_number)
			total := num.add(num2)

			total = total.reduce()

			println(total.mag())

			if total.mag() > max {
				max = total.mag()
			}

		}

	}

	println(max)

}

type number interface {
	mag() int
	String() string
	split() (number, bool)
	explode(int) (number, bool)
	is_hm_n() bool
	get_parent() *snail_number
}

type snail_number struct {
	parent *snail_number
	right  number
	left   number
}

func parse(parent *snail_number, num string) number {

	if len(num) == 1 {
		h_num, _ := strconv.Atoi(num)
		m := &hum_number{parent, h_num}
		return m
	}

	var left, right string

	depth := 0

	collect_left := true

	for _, ch := range num {
		if ch == '[' {
			depth++
			if depth == 1 {
				continue
			}
		} else if ch == ']' {
			if depth == 1 {
				continue
			}
			depth--
		}
		if depth == 1 && ch == ',' {
			collect_left = false
			continue
		}

		if collect_left {
			left += string(ch)
		} else {
			right += string(ch)
		}
	}

	self := snail_number{parent: parent}

	left_num := parse(&self, left)
	right_num := parse(&self, right)

	self.left = left_num
	self.right = right_num

	return &self
}

type hum_number struct {
	parent *snail_number
	value  int
}

func (i hum_number) mag() int {
	return i.value
}

func (sn *snail_number) reduce() *snail_number {

	did_split := true
	did_explode := true

	for did_split || did_explode {
		p, b := sn.explode(0)
		did_explode = b
		sn = p.(*snail_number)
		if did_explode {
		}
		for did_explode {

			p, b := sn.explode(0)
			did_explode = b
			sn = p.(*snail_number)
		}

		p, b = sn.split()
		sn = p.(*snail_number)
		did_split = b

	}

	return sn
}

func (sn *snail_number) add(sn2 *snail_number) *snail_number {

	new_sn := snail_number{}
	sn.parent = &new_sn

	sn2.parent = &new_sn

	new_sn.left = sn
	new_sn.right = sn2

	return &new_sn

}

func (i snail_number) mag() int {

	return 3*i.left.mag() + 2*i.right.mag()
}

func (sn snail_number) String() string {

	return fmt.Sprintf("[%s,%s]", sn.left.String(), sn.right.String())
}

func (hn hum_number) String() string {
	return fmt.Sprint(hn.value)
}

func (sn *snail_number) split() (number, bool) {
	left, split2 := sn.left.split()

	if split2 {
		sn.left = left
		return sn, split2

	}

	right, split1 := sn.right.split()
	sn.left = left
	sn.right = right
	return sn, (split1 || split2)
}

func (hm hum_number) split() (number, bool) {

	if hm.value >= 10 {
		lv := hm.value / 2
		rv := hm.value - lv
		sn := snail_number{parent: hm.parent}
		sn.left = &hum_number{&sn, lv}
		sn.right = &hum_number{&sn, rv}
		return &sn, true
	}
	return &hm, false
}

func (sn *snail_number) explode(depth int) (number, bool) {

	if depth == 4 && sn.left.is_hm_n() && sn.right.is_hm_n() {

		old := sn
		move := old.parent

		for move != nil {

			if move.left.is_hm_n() || *move.left.(*snail_number) != *old {

				descend := move.left

				for !descend.is_hm_n() {
					descend = descend.(*snail_number).right
				}
				descend.(*hum_number).value += sn.left.(*hum_number).value
				break
			}

			old = move
			move = move.parent

		}

		old = sn
		move = old.parent

		for move != nil {

			if move.right.is_hm_n() || *move.right.(*snail_number) != *old {

				descend := move.right

				for !descend.is_hm_n() {
					descend = descend.(*snail_number).left
				}

				descend.(*hum_number).value += sn.right.(*hum_number).value

				break
			}

			old = move
			move = move.parent
		}

		if sn.parent.right.is_hm_n() || sn.parent.right.(*snail_number) != sn {
			sn.parent.left = &hum_number{sn.parent, 0}

		} else {
			sn.parent.right = &hum_number{sn.parent, 0}
		}

		return sn, true

	}

	_, did_explode := sn.left.explode(depth + 1)

	if !did_explode {
		_, did_explode = sn.right.explode(depth + 1)
	}

	return sn, did_explode
}

func (hm hum_number) is_hm_n() bool {
	return true
}

func (sn snail_number) is_hm_n() bool {
	return false
}

func (hm hum_number) explode(depth int) (number, bool) {
	return &hm, false
}

func (sn snail_number) get_parent() *snail_number {
	return sn.parent
}

func (hm hum_number) get_parent() *snail_number {
	return hm.parent
}

func (hm *hum_number) add(num int) {
	hm.value += num
}
