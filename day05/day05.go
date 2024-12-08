package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_ordering(scanner *bufio.Scanner) (map[int][]int, error) {
    ordering_map := make(map[int][]int)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        line_splitted := strings.Split(line, "|")
        if len(line_splitted) != 2 {
            return nil, errors.New("expected two numbers in ordering")
        }

        first_number, err := strconv.Atoi(line_splitted[0])
        if err != nil {
            return nil, err
        }
        second_number, err := strconv.Atoi(line_splitted[1])
        if err != nil {
            return nil, err
        }
        
        ordering_map[first_number] = append(ordering_map[first_number], second_number)
    }
    return ordering_map, nil
}

func read_updates(scanner *bufio.Scanner) ([][]int, error) {
    updates := make([][]int, 0)
    for scanner.Scan() {
        line := scanner.Text()
        line_splitted := strings.Split(line, ",")

        new_updates := make([]int, 0)
        for i := 0; i < len(line_splitted); i++ {
            value, err := strconv.Atoi(line_splitted[i])
            if err != nil {
                return nil, err
            }
            new_updates = append(new_updates, value)
        }
        updates = append(updates, [][]int{new_updates}...)
    }
    return updates, nil
}

func read_input(file_name string) (map[int][]int, [][]int, error) {
    file, err := os.Open(file_name)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    ordering_map, err := read_ordering(scanner)
    if err != nil {
        return nil, nil, err
    }
    updates, err := read_updates(scanner)
    if err != nil {
        return nil, nil, err
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    return ordering_map, updates, nil
}

func part1(ordering_map map[int][]int, updates [][]int) int {
    result_of_middle_numbers := 0
    for _, updates_line := range updates {
        already_used_numbers := make(map[int]bool)
        is_valid := true
        for _, number := range updates_line {
            for _, must_be_behind_number := range ordering_map[number] {
                if already_used_numbers[must_be_behind_number] {
                    is_valid = false
                    break
                }
            }

            if !is_valid {
                break
            }
            already_used_numbers[number] = true
        }

        if is_valid {
            result_of_middle_numbers += updates_line[len(updates_line) / 2]
        }
    }

    return result_of_middle_numbers
}

func part2(ordering_map map[int][]int, updates [][]int) int {
    result_of_middle_numbers := 0
    for _, updates_line := range updates {
        already_used_numbers := make(map[int]bool)

        line := make([]int, len(updates_line))
        copy(line, updates_line)

        is_valid := true
        for i := 0; i < len(line); i++ {
            number := line[i]

            // find the smallest index of a number that is breaking the ordering
            smallest_index := math.MaxInt
            for _, must_be_behind_number := range ordering_map[number] {
                if already_used_numbers[must_be_behind_number] {
                    is_valid = false
                    
                    for j := 0; j < i; j++ {
                        if line[j] == must_be_behind_number {
                            smallest_index = min(j, smallest_index)
                            break
                        }
                    }
                }
            }
            
            if smallest_index != math.MaxInt {
                // We will move Wrong Number just before the smallest index Number that breaks the ordering
                // Other numbers in the checked part of the sequence must be correctly placed, because that
                // part of the sequence is left intact
                prev_num := line[smallest_index]
                line[smallest_index] = line[i]
                for k := smallest_index + 1; k <= i; k++ {
                    next_prev_num := line[k]
                    line[k] = prev_num
                    prev_num = next_prev_num
                }
            }
            already_used_numbers[number] = true
        }

        if !is_valid {
            result_of_middle_numbers += line[len(line) / 2]
        }
    }

    return result_of_middle_numbers
}

func solve(file_name string) error {
    ordering_map, updates, err := read_input(file_name)
    if err != nil {
        return err
    }

    fmt.Println("Part1 result:", part1(ordering_map, updates))
    fmt.Println("Part2 result:", part2(ordering_map, updates))

    return nil
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage day.exe file")
        return
    }

    if err := solve(os.Args[1]); err != nil {
        log.Fatalf("Error: %v", err)
    }
}
