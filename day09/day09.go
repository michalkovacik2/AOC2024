package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
    index int
    free_space int
}

func read_input(file_name string) ([]byte, error) {
    file, err := os.Open(file_name)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var input []byte
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input = scanner.Bytes()
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return input, nil
}

func create_file_system(input []byte, size_of_file_system int) ([]uint16, [][]Node) {
    file_system := make([]uint16, size_of_file_system)
    bins := make([][]Node, 10)
    for i := 0; i < len(bins); i++ {
        bins[i] = make([]Node, 0)
    }

    // IDs in file system are +1 of their real ID (0 means EMPTY)
    var id uint16 = 1
    index := 0
    for i := 0; i < len(input); i++ {
        size := int(input[i] - '0')
        
        if i % 2 == 0 {
            for j := 0; j < size; j++ {
                file_system[index] = id
                index++
            }
            id++
        } else {
            if size != 0 {
                bins[size] = append(bins[size], Node{index: index, free_space: size})
                index += size
            }
        }
    }
    return file_system, bins
}

func calculate_checksum(file_system []uint16) int {
    checksum := 0
    for i := 0; i < len(file_system); i++ {
        if file_system[i] != 0 {
            checksum += i * int(file_system[i] - 1)
        }
    }
    return checksum
}

func find_next_free_index(file_system []uint16, index_free int) int {
    for file_system[index_free] != 0 {
        index_free++
        if index_free >= len(file_system) {
            return -1
        }
    }
    return index_free
}

func find_next_used_index(file_system []uint16, index_used int) int {
    for file_system[index_used] == 0 {
        index_used--
        if index_used < 0 {
            return -1
        }
    }
    return index_used
}

func find_next_used_index_and_size(file_system []uint16, index_used int) (int, int) {
    next_used_index := find_next_used_index(file_system, index_used)
    if next_used_index == -1 {
        return -1, -1
    }
    size := 0
    for i := next_used_index; i > 0 && file_system[i] == file_system[next_used_index]; i-- {
        size++
    }
    return next_used_index, size
}

func repair_bins(bins [][]Node, bin Node) {
    if bin.free_space != 0 {
        bins[bin.free_space] = append([]Node{bin}, bins[bin.free_space]...)
        for i := 1; i < len(bins[bin.free_space]); i++ {
            if bins[bin.free_space][i - 1].index > bins[bin.free_space][i].index {
                tmp := bins[bin.free_space][i - 1]
                bins[bin.free_space][i - 1] = bins[bin.free_space][i]
                bins[bin.free_space][i] = tmp
            }
        }
    }
}

func find_best_bin(bins [][]Node, used_size int) *Node {
    var best_node *Node = nil
    for i := used_size; i < len(bins); i++ {
        if len(bins[i]) > 0 {
            if best_node == nil {
                best_node = &bins[i][0]
            } else if best_node.index > bins[i][0].index {
                best_node = &bins[i][0]
            }
        }
    }
    return best_node
}

func part1(input []byte, size_of_file_system int) int {
    file_system, _ := create_file_system(input, size_of_file_system)
    index_free := find_next_free_index(file_system, 0)
    index_used := find_next_used_index(file_system, len(file_system) - 1)

    for index_free < index_used {
        file_system[index_free] = file_system[index_used]
        file_system[index_used] = 0

        index_free = find_next_free_index(file_system, index_free)
        index_used = find_next_used_index(file_system, index_used)
    }

    return calculate_checksum(file_system)
}

func part2(input []byte, size_of_file_system int) int {
    file_system, bins := create_file_system(input, size_of_file_system)
    index_used, used_size := find_next_used_index_and_size(file_system, len(file_system) - 1)
    
    for index_used > 0 {
        best_node := find_best_bin(bins, used_size)
        
        if best_node != nil && 
        index_used > best_node.index && 
        best_node.free_space >= used_size {
            // Fill the real file system
            idx_free := best_node.index
            idx_used := index_used
            for i := 0; i < used_size; i++ {
                file_system[idx_free] = file_system[idx_used]
                file_system[idx_used] = 0
                idx_free++
                idx_used--
            }

            old_size := best_node.free_space
            best_node.index += used_size
            best_node.free_space -= used_size

            bins[old_size] = bins[old_size][1:]
            if best_node.free_space != 0 {
                repair_bins(bins, *best_node)
            }
        }

        index_used, used_size = find_next_used_index_and_size(file_system, index_used - used_size)
    }

    return calculate_checksum(file_system)
}

func solve(file_name string) error {
    input, err := read_input(file_name)
    if err != nil {
        return err
    }
    
    size_of_filesystem := 0
    for i := 0; i < len(input); i++ {
        size_of_filesystem += int(input[i] - '0');
    }

    fmt.Println("Part1: Checksum:", part1(input, size_of_filesystem))
    fmt.Println("Part2: Checksum:", part2(input, size_of_filesystem))

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
