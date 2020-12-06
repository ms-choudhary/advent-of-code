package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	forest := []string{}
	for scanner.Scan() {
		forest = append(forest, scanner.Text())
	}

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := 1
	for _, s := range slopes {
		tcount := treeCount(forest, s[0], s[1])
		fmt.Println(tcount)
		res *= tcount
	}

	fmt.Println(res)
}

func treeCount(forest []string, right, down int) int {
	count := 0
	for i, j := 0, 0; i < len(forest); i, j = i+down, j+right {
		if '#' == forest[i][j%len(forest[i])] {
			count++
		}
	}
	return count
}
