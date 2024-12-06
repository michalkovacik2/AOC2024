package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func read_input(file_name string) ([]string, error) {
    var input []string
    file, err := os.Open(file_name)
    if err != nil {
        return nil, err
    }
    defer file.Close() // defers waits until return of current function to call itself, pretty neat

    line_size := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(input) == 0 {
            line_size = len(line)
            input = append(input, strings.Repeat(".", line_size + 2))
        }
        input = append(input, "." + line + ".")
    }
    input = append(input, strings.Repeat(".", line_size + 2))

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return input, nil
}

func contains_str(matrix []string, i int, j int, needle string, direction []int) bool {
    index := 0
    
    for matrix[i][j] == needle[index] {
        index++
        i += direction[0]
        j += direction[1]
        
        if index == 4 {
            return true
        }
    }
    return false
}

func part1(matrix []string) {
    directions := [][]int {
        { -1, -1 },
        { -1,  1 },
        { -1,  0 },
        {  0, -1 },
        {  0,  1 }, 
        {  1,  0 },
        {  1,  1 }, 
        {  1, -1 },
    }

    count := 0
    for i := 1; i < len(matrix) - 1; i++ {
        for j := 1; j < len(matrix[i]) - 1; j++ {
            for k := 0; k < len(directions); k++ {
                if contains_str(matrix, i, j, "XMAS", directions[k]){
                    count++
                }
            }
        }
    }

    fmt.Println("Part1 Count:", count)
}

func part2(matrix []string) {
    count := 0
    for i := 1; i < len(matrix) - 1; i++ {
        for j := 1; j < len(matrix[0]) - 1; j++ {
            if matrix[i][j] == 'A' && 
                ((matrix[i - 1][j - 1] == 'M' && matrix[i + 1][j + 1] == 'S') ||
                 (matrix[i - 1][j - 1] == 'S' && matrix[i + 1][j + 1] == 'M')) &&
                ((matrix[i + 1][j - 1] == 'M' && matrix[i - 1][j + 1] == 'S') ||
                 (matrix[i + 1][j - 1] == 'S' && matrix[i - 1][j + 1] == 'M')) {
                    count++
                }
        }
    } 

    fmt.Println("Part2 Count:", count)
}

func solve(file_name string) error {
    matrix, err := read_input(file_name)
    if err != nil {
        return err
    }

    part1(matrix)
    part2(matrix)

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
