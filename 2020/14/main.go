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
	memory := map[uint64]uint64{}
	var mask string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mask") {
			mask = strings.Split(line, "=")[1][1:]
		} else {
			var loc uint64
			var val uint64

			if n, err := fmt.Sscanf(line, "mem[%d] = %d", &loc, &val); err != nil || n != 2 {
				log.Fatalf("failed to parse input %v", err)
			}

			xbits := []uint64{}
			for i, c := range mask {
				if c == '1' {
					loc = setBit(loc, uint64(35-i))
				} else if c == 'X' {
					xbits = append(xbits, uint64(35-i))
				}
			}

			locs := []uint64{}
			for i := uint64(0); i < pow(uint64(2), uint64(len(xbits))); i++ {
				l := loc
				//fmt.Printf("%d = ", i)
				for bi, v := range xbits {
					if bitSet(i, uint64(bi)) {
						l = setBit(l, v)
						//fmt.Printf("%d", 1)
					} else {
						l = clearBit(l, v)
						//fmt.Printf("%d", 0)
					}
				}
				//fmt.Println()
				locs = append(locs, l)
			}

			for _, v := range locs {
				//fmt.Printf("mem[%d] = %d\n", v, val)
				memory[v] = val
			}
		}
	}

	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}

	fmt.Println(sum)
}

func pow(n, base uint64) uint64 {
	p := uint64(1)
	for i := uint64(0); i < base; i++ {
		p *= n
	}
	return p
}

func setBit(n uint64, pos uint64) uint64 {
	n |= (1 << pos)
	return n
}

func bitSet(n uint64, pos uint64) bool {
	if n&(1<<pos) > 0 {
		return true
	}
	return false
}

func clearBit(n uint64, pos uint64) uint64 {
	n &= ^(1 << pos)
	return n
}
