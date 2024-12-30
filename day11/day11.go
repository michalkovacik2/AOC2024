package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_input(file_name string) ([]string, error) {
    var input []string
    file, err := os.Open(file_name)
    if err != nil {
        return nil, err
    }
    defer file.Close() // defers waits until return of current function to call itself, pretty neat

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        input = strings.Split(line, " ")
    }
    
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return input, nil
}

func blink(number string) []string {
    var result []string
    if number == "0" {
        result = append(result, "1")
        return result
    }

    if len(number) % 2 == 0 {
        num1 := number[:len(number) / 2]
        num2 := number[len(number) / 2:]

        num1 = strings.TrimLeft(num1, "0")
        if len(num1) == 0 {
            num1 = "0"
        }
        num2 = strings.TrimLeft(num2, "0")
        if len(num2) == 0 {
            num2 = "0"
        }
        result = append(result, num1, num2)
        return result
    }

    num, _ := strconv.ParseUint(number, 10, 0)
    num = num * 2024
    result = append(result, strconv.FormatUint(num, 10))
    return result
}

func part1(input []string, num_blinks int) int {
    // Too slow for 75 blinks
    current_numbers := input
    for i := 0; i < num_blinks; i++ {
        var new_numbers []string
        for _, number := range current_numbers {
            after_blink_numbers := blink(number)
            new_numbers = append(new_numbers, after_blink_numbers...)
        }
        current_numbers = new_numbers
    }
    return len(current_numbers)
}

func part2(input []string, num_blinks int) int {
    memoization := make([]map[string]int, 75)
    for i := 0; i < len(memoization); i++ {
        memoization[i] = make(map[string]int)
    }

    total := 0
    for _, number := range input {
        total += recursive_solve(number, 1, num_blinks, &memoization)
    }

    return total
}

func recursive_solve(number string, current_blink int, max_blink int, memoization *[]map[string]int) int {
    new_numbers := blink(number)

    if current_blink >= max_blink {
        return len(new_numbers)
    }

    if value, exists := (*memoization)[current_blink][number]; exists {
        return value
    }

    total := 0
    for _, new_number := range new_numbers {
        total += recursive_solve(new_number, current_blink + 1, max_blink, memoization)    
    }
    (*memoization)[current_blink][number] = total

    return total
}

func solve(file_name string) error {
    input, err := read_input(file_name)
    if err != nil {
        return err
    }
    
    fmt.Println("Part1: After 25 blinks:", part1(input, 25))
    fmt.Println("Part2: After 75 blinks:", part2(input, 75))
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
