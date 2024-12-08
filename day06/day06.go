package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
    UP    = 0b0001
    RIGHT = 0b0010
    DOWN  = 0b0100
    LEFT  = 0b1000
)

func read_input(file_name string) ([][]byte, int, int, error) {
    var input [][]byte
    file, err := os.Open(file_name)
    if err != nil {
        return nil, -1, -1, err
    }
    defer file.Close()

    line_size := 0
    current_x := 0
    guard_x, guard_y := 0, 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(input) == 0 {
            line_size = len(line)
            input = append(input, []byte(strings.Repeat("0", line_size + 2)))
            current_x++
        }
        input = append(input, []byte("0" + line + "0"))
        if idx := strings.Index(string(input[len(input) - 1]), "^"); idx != -1 {
            guard_x = current_x
            guard_y = idx
        }
        current_x++
    }
    input = append(input, []byte(strings.Repeat("0", line_size + 2)))

    if err := scanner.Err(); err != nil {
        return nil, -1, -1, err
    }

    return input, guard_x, guard_y, nil
}

func part1(matrix [][]byte, guard_x int, guard_y int, direction_index int) (int, bool) {
    visited := make([][]int, len(matrix))
    for i := range visited {
        visited[i] = make([]int, len(matrix[0]))
    }
    
    switch direction_index {
        case 0: visited[guard_x][guard_y] = UP
        case 1: visited[guard_x][guard_y] = RIGHT
        case 2: visited[guard_x][guard_y] = DOWN
        case 3: visited[guard_x][guard_y] = LEFT
    }
    num_visited := 1

    directions := [][]int {
        {-1,  0 },
        { 0,  1 },
        { 1,  0 },
        { 0, -1 },
    }

    x, y := guard_x, guard_y
    for {
        new_x := x + directions[direction_index][0]
        new_y := y + directions[direction_index][1]
        
        if matrix[new_x][new_y] == '0' {
            break
        }

        if matrix[new_x][new_y] == '#' {
            direction_index = (direction_index + 1) % len(directions)
            continue
        }

        if visited[new_x][new_y] == 0 {
            num_visited++
        }
        
        // Loop detection
        if (direction_index == 0 && ((visited[new_x][new_y] & UP) != 0)) ||
           (direction_index == 1 && ((visited[new_x][new_y] & RIGHT) != 0)) ||
           (direction_index == 2 && ((visited[new_x][new_y] & DOWN) != 0)) ||
           (direction_index == 3 && ((visited[new_x][new_y] & LEFT) != 0)) {
                return -1, true
            }

        switch direction_index {
            case 0: visited[new_x][new_y] |= UP
            case 1: visited[new_x][new_y] |= RIGHT 
            case 2: visited[new_x][new_y] |= DOWN 
            case 3: visited[new_x][new_y] |= LEFT 
        }

        x = new_x
        y = new_y
    }
    return num_visited, false
}

func part2(matrix [][]byte, guard_x int, guard_y int) int {
    direction_index := 0
    directions := [][]int {
        {-1,  0 },
        { 0,  1 },
        { 1,  0 },
        { 0, -1 },
    }

    obstruction_placed := make([][]bool, len(matrix))
    for i := range obstruction_placed {
        obstruction_placed[i] = make([]bool, len(matrix[0]))
    }

    loop_counter := 0
    x, y := guard_x, guard_y
    for {
        new_x := x + directions[direction_index][0]
        new_y := y + directions[direction_index][1]
        
        if matrix[new_x][new_y] == '0' {
            break
        }

        if matrix[new_x][new_y] == '#' {
            direction_index = (direction_index + 1) % len(directions)
            continue
        }

        if !(new_x == guard_x && new_y == guard_y) && !obstruction_placed[new_x][new_y] {
            // Try putting # in front of current position
            matrix[new_x][new_y] = '#'
            obstruction_placed[new_x][new_y] = true

            _, loop_detected := part1(matrix, x, y, direction_index)
            if loop_detected {
                loop_counter++
            }

            // Revert the change
            matrix[new_x][new_y] = '.'
        }

        x = new_x
        y = new_y
    }
    return loop_counter
}

func solve(file_name string) error {
    matrix, guard_x, guard_y, err := read_input(file_name)
    if err != nil {
        return err
    }

    visited, _ := part1(matrix, guard_x, guard_y, 0) // Start at UP
    fmt.Println("Part1 visited:", visited)
    fmt.Println("Part2 loops:  ", part2(matrix, guard_x, guard_y))

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
