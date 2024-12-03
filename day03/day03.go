package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_input(file_name string) (string, error) {
    input_text := ""
    file, err := os.Open(file_name)
    if err != nil {
        return "", err
    }
    defer file.Close() // defers waits until return of current function to call itself, pretty neat

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input_text += scanner.Text()
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return input_text, nil
}

func part1(input_text string) error {
    regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
    regex_result := regex.FindAll([]byte(input_text), -1)

    result := 0
    for _, item := range regex_result {
        mul_expression := string(item)
        mul_expression = mul_expression[4:len(mul_expression)-1]
        splitted := strings.Split(mul_expression, ",")
        num1, err := strconv.Atoi(splitted[0])
        if err != nil {
            return err
        }
        num2, err := strconv.Atoi(splitted[1])
        if err != nil {
            return err
        }
        result += num1 * num2
    }

    fmt.Println("Part1 result:", result)
    return nil
}

func part2(input_text string) error {
    regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|(do\(\))|(don't\(\))`)
    regex_result := regex.FindAll([]byte(input_text), -1)

    result := 0
    mul_is_enabled := true
    for _, item := range regex_result {
        str := string(item)
        if str == "do()" {
            mul_is_enabled = true
        } else if str == "don't()" {
            mul_is_enabled = false
        } else {
            if mul_is_enabled {
                str = str[4:len(str)-1]
                splitted := strings.Split(str, ",")
                num1, err := strconv.Atoi(splitted[0])
                if err != nil {
                    return err
                }
                num2, err := strconv.Atoi(splitted[1])
                if err != nil {
                    return err
                }
                result += num1 * num2
            }
        }
    }

    fmt.Println("Part2 result:", result)
    return nil
}

func solve(file_name string) error {
    input_text, err := read_input(file_name)
    if err != nil {
        return err
    }

    if err := part1(input_text); err != nil {
        return err
    }

    if err := part2(input_text); err != nil {
        return err
    }

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
