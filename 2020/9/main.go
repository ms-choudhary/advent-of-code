package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const preambleLen = 25

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	nos := []int{}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nos = append(nos, n)
	}

	prevGroup := map[int]bool{}

	for i := 0; i < preambleLen; i++ {
		prevGroup[nos[i]] = true
	}

	for i := preambleLen; i < len(nos); i++ {
		sumexists := false
		for j := i - preambleLen; j < i; j++ {
			if _, ok := prevGroup[nos[i]-nos[j]]; ok {
				sumexists = true
				break
			}
		}

		if !sumexists {
			fmt.Println(nos[i])
			return
		}

		prevGroup[nos[i]] = true
		delete(prevGroup, nos[i-preambleLen])
	}
}
