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
	groupAns := ""
	sumOfCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sumOfCount += countUniqAns(groupAns)
			groupAns = ""
		} else if groupAns == "" {
			groupAns = line
		} else {
			groupAns += line
		}
	}

	sumOfCount += countUniqAns(groupAns)

	fmt.Println(sumOfCount)
}

func countUniqAns(ans string) int {
	uniqAns := map[rune]bool{}

	for _, c := range ans {
		uniqAns[c] = true
	}

	return len(uniqAns)
}
