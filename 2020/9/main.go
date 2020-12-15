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
	nos := []int{}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nos = append(nos, n)
	}

	//fmt.Println(findEncryptionWeakness(nos, 257342611))
	arr := findEncryptionWeakness(nos, 257342611)
	sort.Ints(arr)
	fmt.Println(arr[0] + arr[len(arr)-1])
}

func findEncryptionWeakness(nos []int, n int) []int {
	i := 0
	j := 0
	cursum := 0
	for j < len(nos) {

		if cursum < n {
			cursum += nos[j]
			j++
		}

		if cursum == n {
			return nos[i:j]
		}

		if cursum > n {
			cursum -= nos[i]
			i++
		}
	}
	return []int{}
}
