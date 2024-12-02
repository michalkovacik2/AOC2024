package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func read_input(file_name string) ([][]int, error) {
    lines := make([][]int, 0)

    file, err := os.Open(file_name)
    if err != nil {
        return nil, err
    }
    defer file.Close() // defers waits until return of current function to call itself, pretty neat

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        line_parts := strings.Fields(line)

        new_line := make([]int, 0)
        for _, line_part := range line_parts {
            val, err := strconv.Atoi(line_part)
            if err != nil {
                return nil, err
            }
            new_line = append(new_line, val)
        }
        lines = append(lines, [][]int{new_line}...)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

func is_safe(line []int) bool {
    var is_increasing bool
    if len(line) < 2 {
        return false // Not sure
    }

    if line[0] > line[1] {
        is_increasing = false
    } else if line[0] < line[1] {
        is_increasing = true
    } else {
        return false
    }

    for i := 0; i < len(line) - 1; i++ {
        if is_increasing {
            if line[i + 1] < line[i] {
                return false
            } 
        } else {
            if line[i + 1] > line[i] {
                return false
            }
        }

        diff := line[i + 1] - line[i]
        if diff < 0 {
            diff = -diff
        }

        if diff < 1 || diff > 3 {
            return false
        }
    }

    return true
}

func part1_and_part2(file_name string) error {
    lines, err := read_input(file_name)
    if err != nil {
        return err
    }

    num_safe_reportsV1 := 0
    num_safe_reportsV2 := 0

    for _, line := range lines {
        if is_safe(line) {
            num_safe_reportsV1++
            num_safe_reportsV2++
        } else {
            // Try erasing every single element and check if it became safe :(
            // N^2 not great not terrible (the lines are very small)
            // You can probably optimize this by returning the index where it becomes bad 
            // and only try removing elements around it.
            for i := 0; i < len(line); i++ {
                if is_safe(slices.Concat(line[:i], line[i+1:])) {
                    num_safe_reportsV2++
                    break
                }
            }
        }
    }

    fmt.Printf("Number of safe reports version 1: %d\n", num_safe_reportsV1)
    fmt.Printf("Number of safe reports version 2: %d\n", num_safe_reportsV2)
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
