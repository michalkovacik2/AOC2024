package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var DIRECTIONS = [][]int {
    { -1,  0 },
    {  0,  1 },
    {  1,  0 },
    {  0, -1 },
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

func dfs(matrix [][]byte,  visited [][]bool, x int, y int, score *int) {
    if visited != nil {
        visited[x][y] = true
    }

    current_value := matrix[x][y]
    if current_value == '9' {
        (*score)++
        return
    }
    
    for _, direction := range DIRECTIONS {
        new_x, new_y := x + direction[0], y + direction[1]
        if matrix[new_x][new_y] - current_value == 1 && 
           (visited == nil || !visited[new_x][new_y]) {
            dfs(matrix, visited, new_x, new_y, score)
        }
    }
}

func solve_parts(matrix [][]byte, track_visited bool) int {
    var score int = 0
    for i := 1; i < len(matrix) - 1; i++ {
        for j := 1; j < len(matrix[i]) - 1; j++ {
            if matrix[i][j] == '0' {
                var visited [][]bool = nil
                if track_visited {
                    visited = make([][]bool, len(matrix))
                    for i := 0; i < len(matrix); i++ {
                        visited[i] = make([]bool, len(matrix[i]))
                    }
                }

                dfs(matrix, visited, i, j, &score)
            }
        }
    }
    return score
}

func solve(file_name string) error {
    matrix, err := read_matrix(file_name, ".")
    if err != nil {
        return err
    }
    
    fmt.Println("Part1: Sum of scores of trailheads:", solve_parts(matrix, true))
    fmt.Println("Part2: Sum of ratings of trailheads:", solve_parts(matrix, false))
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
