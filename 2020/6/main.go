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

	defer f.Close()

	scanner := bufio.NewScanner(f)
	groupAns := []string{}
	sumOfCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sumOfCount += countCommonAns(groupAns)
			groupAns = []string{}
		} else {
			groupAns = append(groupAns, line)
		}
	}

	sumOfCount += countCommonAns(groupAns)

	fmt.Println(sumOfCount)
}

func countCommonAns(ans []string) int {
	commonAns := map[rune]int{}

	for _, a := range ans {
		for _, c := range a {
			commonAns[c]++
		}
	}

	count := 0
	for _, v := range commonAns {
		if v == len(ans) {
			count++
		}
	}
	return count
}
