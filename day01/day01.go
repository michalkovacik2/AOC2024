package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_input(file_name string) ([]int, []int, error) {
	var list1, list2 []int

	file, err := os.Open(file_name)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close() // defers waits until return of current function to call itself, pretty neat

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_parts := strings.Fields(line)

		id1, err := strconv.Atoi(line_parts[0])
		if err != nil {
			return nil, nil, err
		}
		id2, err := strconv.Atoi(line_parts[1])
		if err != nil {
			return nil, nil, err
		}

		list1 = append(list1, id1)
		list2 = append(list2, id2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil
}

func calculateDistance(list1, list2 []int) uint64 {
	var distance uint64 = 0
	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		distance += uint64(diff)
	}
	return distance
}

func calculateSimilarity(list1, list2 []int) uint64 {
	var similarity uint64 = 0
	idx1, idx2 := 0, 0
	for idx1 < len(list1) {
		current_num := list1[idx1]
		duplicates_list2 := 0

		for idx2 < len(list2) && list2[idx2] <= current_num {
			if list2[idx2] == current_num {
				duplicates_list2++
			}
			idx2++
		}

		duplicates_list1 := 1
		idx1++
		for idx1 < len(list1) && list1[idx1] == current_num {
			duplicates_list1++
			idx1++
		}
		similarity += uint64(current_num * duplicates_list1 * duplicates_list2)
	}
	return similarity
}

func part1_and_part2(file_name string) error {
	list1, list2, err := read_input(file_name)
	if err != nil {
		return err
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	// Part 1
	distance := calculateDistance(list1, list2)
	// Part 2
	similarity := calculateSimilarity(list1, list2)

	fmt.Println("Total distance:  ", distance)
	fmt.Println("Similarity score:", similarity)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage day.exe file")
		return
	}

	if err := part1_and_part2(os.Args[1]); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
