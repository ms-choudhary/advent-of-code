package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	seats := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		row := decodeSeat(s[:len(s)-3], 0, 127)
		col := decodeSeat(s[len(s)-3:], 0, 7)

		sid := (row * 8) + col
		seats = append(seats, sid)
	}

	sort.Ints(seats)

	fmt.Println(findMissingNo(seats))
}

func findMissingNo(list []int) int {
	n := list[0]
	for _, v := range list {
		if v == n+1 {
			return n
		}
		n++
	}
	return -1
}

func decodeSeat(s string, low, high int) int {
	for _, c := range s {
		//fmt.Printf("low: %v, high: %v\n", low, high)
		if c == 'F' || c == 'L' {
			low, high = findRange(true, low, high)
		} else if c == 'B' || c == 'R' {
			low, high = findRange(false, low, high)
		}
	}
	return low
}

func findRange(lower bool, low, high int) (int, int) {
	mid := low + ((high - low) / 2)

	if lower {
		return low, mid
	}
	return mid + 1, high
}
