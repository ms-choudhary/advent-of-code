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
	ratings := []int{0}
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			ratings = append(ratings, i)
		}
	}

	sort.Ints(ratings)

	ratings = append(ratings, ratings[len(ratings)-1]+3)

	irrelevant := 0
	for i := 0; i < len(ratings)-1; {
		j := i + 1
		for ; j < len(ratings); j++ {
			if ratings[j]-ratings[i] > 3 {
				break
			}
			irrelevant++
		}
		i = j - 1
		irrelevant -= 1
		fmt.Printf("%v %v\n", ratings[i], irrelevant)
	}

	res := 1
	for i := 0; i < irrelevant; i++ {
		res *= 2
	}

	fmt.Println(res)
}
