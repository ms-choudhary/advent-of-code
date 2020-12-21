package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	memory := map[int]uint64{}
	var mask string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, "=")[1][1:]
		} else {
			var loc int
			var val uint64

			if n, err := fmt.Sscanf(line, "mem[%d] = %d", &loc, &val); err != nil || n != 2 {
				log.Fatalf("failed to parse input %v", err)
			}

			for i, c := range mask {
				if c == '1' {
					val |= (1 << uint64(35-i))
				} else if c == '0' {
					val &= ^(1 << uint64(35-i))
				}
			}

			//fmt.Printf("mem[%d] = %d\n", loc, val)
			memory[loc] = val
		}
	}

	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}

	fmt.Println(sum)
}
