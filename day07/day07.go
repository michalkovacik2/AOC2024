package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func read_input(file_name string) ([][]int ,error) {
    var input [][]int
    file, err := os.Open(file_name)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        splitted := strings.Fields(line)
        line_nums := make([]int, 0)
        for i, text := range splitted {
            if i == 0 {
                text = strings.Split(text, ":")[0]
            } 

            num, err := strconv.Atoi(text)
            if err != nil {
                return nil, err
            }
            line_nums = append(line_nums, num)
        }
        input = append(input, [][]int{line_nums}...)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return input, nil
}

func calculate_result(values []int, current_permutation []int, OP *[]func(int, int) int) int {
    res := values[0]
    for i := 0; i < len(current_permutation); i++ {
        res = (*OP)[current_permutation[i]](res, values[i + 1])
    }
    return res
}

func calculate_permutations(lines [][]int, OP *[]func(int, int) int) int {
    sum := 0
    for _, line := range lines {
        expected_result := line[0]
        values := line[1:]

        current_permutation := make([]int, len(values) - 1)

        run := true
        for run {
            current_result := calculate_result(values, current_permutation, OP)
            if current_result == expected_result {
                sum += expected_result
                break
            }

            for i := len(current_permutation) - 1; i >= 0; i-- {
                current_permutation[i]++
                if current_permutation[i] < len(*OP) {
                    break
                }
                current_permutation[i] = 0
                if i == 0 {
                    run = false
                }
            }
        }
    }
    return sum
}

func solve(file_name string) error {
    lines, err := read_input(file_name)
    if err != nil {
        return err
    }

    OPERATORS1 := []func(int, int) int {
        func (a, b int) int { return a + b },
        func (a, b int) int { return a * b },
    }

    OPERATORS2 := []func(int, int) int {
        func (a, b int) int { return a + b },
        func (a, b int) int { return a * b },
        func (a, b int) int { val, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b)); return val },
    }

    fmt.Println("Part1: Total calibration result:", calculate_permutations(lines, &OPERATORS1))
    fmt.Println("Part2: Total calibration result:", calculate_permutations(lines, &OPERATORS2))
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
