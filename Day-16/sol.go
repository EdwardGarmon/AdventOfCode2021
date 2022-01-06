package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

var hex_map map[string]string

var version_sum int64
var fp int64
var sp int64

func main() {
	hex_map = map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	message := string(file)
	version_sum = 0
	packet, val9 := read_packet(hexToBin(message))
	println(packet, val9, version_sum)

}

func read_packet(packet string) (string, int64) {

	stack := []int64{}

	version := binToDec(packet[:3])
	packet = packet[3:]

	version_sum += version

	id := binToDec(packet[:3])
	packet = packet[3:]

	if id == 4 {

		literal, p := read_literal(packet)
		packet = p

		return packet, literal

	} else {

		length_id := packet[0]
		packet = packet[1:]

		if length_id == '0' {
			pack_num := binToDec(packet[:15])
			packet := packet[15:]

			old_len := int64(len(packet))
			for old_len-int64(len(packet)) != pack_num {
				p, val := read_packet(packet)
				packet = p
				stack = append(stack, val)
			}
			return packet, perform_op(id, stack)

		} else if length_id == '1' {
			pack_num := binToDec(packet[:11])
			packet = packet[11:]

			for i := int64(0); i < pack_num; i++ {
				p, val := read_packet(packet)
				packet = p
				stack = append(stack, val)
			}

			return packet, perform_op(id, stack)
		}
	}

	return "", 0

}

func perform_op(op_code int64, stack []int64) int64 {

	if op_code == 0 {
		return sum(stack)
	} else if op_code == 1 {
		return prod(stack)
	} else if op_code == 2 {
		return find_min(stack)
	} else if op_code == 3 {
		return find_max(stack)
	} else if op_code == 5 {
		return gt(stack)
	} else if op_code == 6 {
		return lt(stack)
	} else if op_code == 7 {
		return et(stack)
	}

	return 0

}

func find_min(stack []int64) int64 {
	min := stack[0]
	for x := 0; x < len(stack); x++ {
		val := stack[x]
		if val < min {
			min = val
		}
	}
	return min
}

func find_max(stack []int64) int64 {
	max := stack[0]
	for x := 0; x < len(stack); x++ {
		val := stack[x]
		if val > max {
			max = val
		}
	}
	return max
}

func sum(stack []int64) int64 {

	sum := int64(0)
	for x := 0; x < len(stack); x++ {
		val := stack[x]
		sum += val
	}
	return sum

}

func prod(stack []int64) int64 {
	prod := int64(1)
	for x := 0; x < len(stack); x++ {
		val := stack[x]
		prod *= val
	}
	return prod
}

func gt(stack []int64) int64 {
	if stack[0] > stack[1] {
		return 1
	}
	return 0
}

func lt(stack []int64) int64 {
	if stack[0] < stack[1] {
		return 1
	}
	return 0

}

func et(stack []int64) int64 {
	if stack[0] == stack[1] {
		return 1
	}
	return 0

}

func read_literal(packet string) (int64, string) {
	lit := ""
	for packet[0] != '0' {
		lit += packet[1:5]
		packet = packet[5:]
	}
	last_chunk := packet[0:5]
	lit += last_chunk[1:]
	packet = packet[5:]

	v := binToDec(lit)
	println(v)
	return v, packet
}

func binToDec(bin string) int64 {
	val, _ := strconv.ParseInt(bin, 2, 64)
	return int64(val)
}

func hexToBin(hex string) string {
	bin := ""
	for _, ch := range hex {
		bin += hex_map[string(ch)]
	}
	return bin
}
