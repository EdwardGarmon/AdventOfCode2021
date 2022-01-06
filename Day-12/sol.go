package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	name      string
	neighbors []string
}

func isSmall(name string) bool {

	return name == strings.ToLower(name)
}

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	cave := make(map[string]Node)

	insert_edge := func(start string, end string) {
		node, exists := cave[start]

		if exists {
			new_ns := append(node.neighbors, end)
			new_node := Node{name: start, neighbors: new_ns}
			cave[start] = new_node

		} else {
			new_node := Node{name: start, neighbors: []string{end}}
			cave[start] = new_node

		}
	}

	for _, edge := range strings.Split(string(file), "\n") {

		conns := strings.Split(edge, "-")

		start := conns[0]
		end := conns[1]

		insert_edge(start, end)
		if start != "start" {
			insert_edge(end, start)
		}

	}

	paths := 0

	var dfs func(string, map[string]int)

	dfs = func(name string, visited map[string]int) {
		for _, neighm := range cave[name].neighbors {
			count, exists := visited[neighm]

			may_visit := true

			if !exists {
				count = 0
			}

			if count == 1 {
				for _, val := range visited {
					if val == 2 {
						may_visit = false
					}
				}
			}

			if neighm == "start" {
				continue
			}

			if count > 1 {
				continue
			}

			if may_visit {

				if neighm == "end" {
					// for key, _ := range visited {
					// 	fmt.Print(key + ", ")
					// }
					// fmt.Print("end")
					// fmt.Println()
					paths++

				} else {
					copy_visited := make(map[string]int)

					for key, value := range visited {
						copy_visited[key] = value
					}

					if isSmall(neighm) {
						copy_visited[neighm] = count + 1
					}

					dfs(neighm, copy_visited)
				}
			}
		}
	}

	// for _, value := range cave {
	// 	println(value.name)
	// 	fmt.Println(strings.Join(value.neighbors, ","))

	// }

	visited := make(map[string]int)

	dfs("start", visited)

	fmt.Printf("this many paths %d \n", paths)

}
