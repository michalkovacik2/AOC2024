package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
    row int
    col int
}

func read_matrix(file_name string, padding_char string) ([][]byte ,error) {
    file, err := os.Open(file_name)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var input [][]byte
    var padding []byte

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(input) == 0 {
            if padding_char != "" {
                padding = []byte(strings.Repeat(padding_char, len(line) + 2))
                input = append(input, padding)
            }
        }
        input = append(input, []byte(padding_char + line + padding_char))
    }

    if padding_char != "" {
        input = append(input, padding)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return input, nil
}

func get_antena_positions(matrix [][]byte) map[byte][]Coord {
    antenas := make(map[byte][]Coord)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] != '.' {
                key := matrix[i][j]
                new_antena_position := Coord {i, j}
                antenas[matrix[i][j]] = append(antenas[key], new_antena_position)
            }
        }
    }
    return antenas
}

func is_in_bounds(matrix [][]byte, coord *Coord) bool {
    return coord.row >= 0 && coord.row < len(matrix) && coord.col >= 0 && coord.col < len(matrix[0])
}

func print_matrix(matrix [][]byte) {
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("%c", matrix[i][j])
        }
        fmt.Print("\n")
    }
    fmt.Print("\n")
}

func deep_copy_matrix(matrix [][]byte) [][]byte {
    matrix_copy := make([][]byte, len(matrix))
    for i := 0; i < len(matrix); i++ {
        matrix_copy[i] = make([]byte, len(matrix[i]))
        copy(matrix_copy[i], matrix[i])
    }
    return matrix_copy
}

func part1(matrix [][]byte, antenas map[byte][]Coord) int {
    number_of_antinodes := 0
    for _, positions := range antenas {
        for i := 0; i < len(positions); i++ {
            for j := i + 1; j < len(positions); j++ {
                diff := Coord{positions[j].row - positions[i].row, positions[j].col - positions[i].col}
                antinode1 := Coord{positions[i].row - diff.row, positions[i].col - diff.col}
                antinode2 := Coord{positions[j].row + diff.row, positions[j].col + diff.col}

                if is_in_bounds(matrix, &antinode1) &&
                matrix[antinode1.row][antinode1.col] != '#' {
                    number_of_antinodes++
                    matrix[antinode1.row][antinode1.col] = '#'
                }
                
                if is_in_bounds(matrix, &antinode2) &&
                matrix[antinode2.row][antinode2.col] != '#' {
                    number_of_antinodes++
                    matrix[antinode2.row][antinode2.col] = '#'
                }
            }
        }
    }
    return number_of_antinodes
}

func part2(matrix [][]byte, antenas map[byte][]Coord) int {
    number_of_antinodes := 0
    for _, positions := range antenas {
        for i := 0; i < len(positions); i++ {
            for j := i + 1; j < len(positions); j++ {
                diff := Coord{positions[j].row - positions[i].row, positions[j].col - positions[i].col}
                
                k := 0
                antinode1 := Coord{positions[i].row - k * diff.row, positions[i].col - k * diff.col}
                for is_in_bounds(matrix, &antinode1) {
                    if matrix[antinode1.row][antinode1.col] != '#' {
                        number_of_antinodes++
                        matrix[antinode1.row][antinode1.col] = '#'
                    }
                    k++
                    antinode1 = Coord{positions[i].row - k * diff.row, positions[i].col - k * diff.col}
                }
                
                k = 0
                antinode2 := Coord{positions[j].row + k * diff.row, positions[j].col + k * diff.col}
                for is_in_bounds(matrix, &antinode2) {
                    if matrix[antinode2.row][antinode2.col] != '#' {
                        number_of_antinodes++
                        matrix[antinode2.row][antinode2.col] = '#'
                    }
                    k++
                    antinode2 = Coord{positions[j].row + k * diff.row, positions[j].col + k * diff.col}
                }
            }
        }
    }
    return number_of_antinodes
}

func solve(file_name string) error {
    matrix, err := read_matrix(file_name, "")
    if err != nil {
        return err
    }
    antenas := get_antena_positions(matrix)

    fmt.Println("Part1: Number of antinodes:", part1(deep_copy_matrix(matrix), antenas))
    fmt.Println("Part2: Number of antinodes:", part2(matrix, antenas))
    // print_matrix(matrix)
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
