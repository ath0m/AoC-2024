package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) *[]int {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic("cannot read file")
	}
	content := strings.TrimSpace(string(file))
	diskmap := []int{}
	for _, ch := range content {
		diskmap = append(diskmap, int(ch-'0'))
	}
	return &diskmap
}

func expandDiskmap(diskmap *[]int) *[]int {
	blocks := []int{}
	for i, size := range *diskmap {
		var id int
		if i%2 == 0 {
			id = i / 2
		} else {
			id = -1
		}
		for j := 0; j < size; j++ {
			blocks = append(blocks, id)
		}
	}
	return &blocks
}

func p1(diskmap *[]int) int {
	blocks := expandDiskmap(diskmap)

	l, r := 0, len(*blocks)-1
	for l < r {
		for l < len(*blocks) && (*blocks)[l] >= 0 {
			l += 1
		}
		for r >= 0 && (*blocks)[r] < 0 {
			r -= 1
		}
		if l < r {
			(*blocks)[l], (*blocks)[r] = (*blocks)[r], (*blocks)[l]
		}
	}

	result := 0
	for i, val := range *blocks {
		if val >= 0 {
			result += i * val
		}
	}

	return result
}

type Block struct {
	id, offset, size int
}

func printBlocks(fileBlocks *[]Block, freeBlocks *[]Block) {
	size := 0

	if len(*fileBlocks) > 0 {
		last := (*fileBlocks)[len(*fileBlocks)-1]
		size = max(size, last.offset+last.size)
	}
	if len(*freeBlocks) > 0 {
		last := (*freeBlocks)[len(*freeBlocks)-1]
		size = max(size, last.offset+last.size)
	}

	mem := []string{}
	for i := 0; i < size; i++ {
		mem = append(mem, "x")
	}

	for _, block := range *fileBlocks {
		for i := block.offset; i < block.offset+block.size; i++ {
			mem[i] = strconv.Itoa(block.id)
		}
	}
	for _, block := range *freeBlocks {
		for i := block.offset; i < block.offset+block.size; i++ {
			mem[i] = "."
		}
	}

	fmt.Println(strings.Join(mem, ""))
}

func p2(diskmap *[]int) int {
	freeBlocks := []Block{}
	fileBlocks := []Block{}

	offset := 0
	for i, size := range *diskmap {
		if i%2 == 0 {
			fileBlocks = append(fileBlocks, Block{id: i / 2, offset: offset, size: size})
		} else {
			freeBlocks = append(freeBlocks, Block{id: -1, offset: offset, size: size})
		}
		offset += size
	}

	// printBlocks(&fileBlocks, &freeBlocks)

	for i := range fileBlocks {
		file := &fileBlocks[len(fileBlocks)-1-i]
		for j := range freeBlocks {
			free := &freeBlocks[j]
			if file.offset < free.offset {
				break
			}
			if file.size <= free.size {
				file.offset = free.offset
				free.offset += file.size
				free.size -= file.size
				break
			}
		}
		// printBlocks(&fileBlocks, &freeBlocks)
	}

	// printBlocks(&fileBlocks, &freeBlocks)

	result := 0
	for _, block := range fileBlocks {
		result += block.id * (block.size*block.offset + (block.size*(block.size-1))/2)
	}

	return result
}

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("argument is required")
	}
	diskmap := parse("input.txt")
	switch args[1] {
	case "p1":
		fmt.Println("p1:", p1(diskmap))
	case "p2":
		fmt.Println("p2:", p2(diskmap))
	}
}
