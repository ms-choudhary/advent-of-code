package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sum = 2020

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var nos []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("not a number %v", err)
		}
		nos = append(nos, n)
	}

	x1, x2, err := twoNsWhichSumTo(sum, nos)
	if err == nil {
		fmt.Printf("1: %v, 2: %v, 1*2: %v\n", x1, x2, x1*x2)
	}

	for i, n1 := range nos {
		n2, n3, err := twoNsWhichSumTo(sum-n1, nos[i:])
		if err == nil {
			fmt.Printf("1: %v, 2: %v, 3: %v, 1*2*3: %v\n", n1, n2, n3, n1*n2*n3)
		}
	}
}

func twoNsWhichSumTo(sum int, nos []int) (int, int, error) {
	cache := map[int]bool{}
	for _, n := range nos {
		if _, ok := cache[sum-n]; ok {
			return n, sum - n, nil
		} else {
			cache[n] = true
		}
	}

	return 0, 0, fmt.Errorf("not found")
}
