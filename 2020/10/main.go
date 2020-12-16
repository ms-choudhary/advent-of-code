package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	ratings := []int{}
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			ratings = append(ratings, i)
		}
	}

	sort.Ints(ratings)

	no1s, no3s := 0, 0
	n := 0
	for i := 0; i < len(ratings); i++ {
		if ratings[i]-n > 3 {
			fmt.Printf("diff betn %v and %v greater than 3\n", ratings[i], n)
			break
		} else if ratings[i]-n == 3 {
			no3s++
		} else if ratings[i]-n == 1 {
			no1s++
		}
		n = ratings[i]
	}

	fmt.Println(no1s * (no3s + 1))
}
