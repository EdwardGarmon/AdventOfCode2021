package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("input file not found")
		return
	}

	start_rules := strings.Split(string(file), "\n\n")

	rules := strings.Split(start_rules[1], "\n")

	rules_map := make(map[string]string)

	for _, rule := range rules {
		var pattern string
		var gen string
		fmt.Sscanf(rule, "%s -> %s", &pattern, &gen)
		rules_map[pattern] = gen
	}

	//initialize the counts map

	cache := make(map[string]map[int]map[string]int)

	var expand_pair func(string, int, int) map[string]int

	expand_pair = func(pair string, depth int, max_depth int) map[string]int {

		counts_map := make(map[string]int)

		if depth == max_depth {
			return counts_map
		}

		pair_map, pair_exists := cache[pair]

		if pair_exists {
			cache_val, cache_hit := pair_map[depth]
			if cache_hit {
				return cache_val
			}
		}

		gen, exists := rules_map[pair]

		if exists {
			incrementMap(gen, counts_map)
			counts_map = combine_map(counts_map, expand_pair(string(pair[0])+gen, depth+1, max_depth))
			counts_map = combine_map(counts_map, expand_pair(gen+string(pair[1]), depth+1, max_depth))
		}

		cache_pair, pair_cache_exists := cache[pair]

		if !pair_cache_exists {
			cache[pair] = make(map[int]map[string]int)
			cache[pair][depth] = counts_map
		} else {
			cache_pair[depth] = counts_map
		}

		return counts_map
	}

	polymer := start_rules[0]
	final_counts := make(map[string]int)
	for _, ch := range polymer {
		incrementMap(string(ch), final_counts)
	}

	for x := 0; x < len(polymer)-1; x += 1 {

		pair := polymer[x : x+2]

		final_counts = combine_map(final_counts, expand_pair(pair, 0, 40))

	}
	small := final_counts[string(polymer[0])]
	large := final_counts[string(polymer[0])]
	for key, value := range final_counts {
		println(key, value)
		if value < small {
			small = value
		}
		if value > large {
			large = value
		}
	}

	println(large - small)
}

func incrementMap(ch string, count_map map[string]int) {

	count, exists := count_map[ch]

	if exists {
		count_map[ch] = count + 1
	} else {
		count_map[ch] = 1
	}

}

func combine_map(map1 map[string]int, map2 map[string]int) map[string]int {
	for key, value := range map2 {
		_, exists := map1[key]
		if exists {
			map1[key] += value
		} else {
			map1[key] = value
		}
	}
	return map1
}
