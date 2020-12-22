package main

import "fmt"

const lastTurn = 30000000

func main() {
	lastSpoken := map[int]int{}
	startingNos := []int{2, 0, 1, 7, 4, 14, 18}

	prevN := startingNos[0]
	turn := 1
	for ; turn < lastTurn; turn++ {
		//fmt.Println(prevN)
		if turn < len(startingNos) {
			lastSpoken[prevN] = turn - 1
			prevN = startingNos[turn]
		} else {
			if _, ok := lastSpoken[prevN]; !ok {
				lastSpoken[prevN] = turn - 1
				prevN = 0
			} else {
				prevSpoken := lastSpoken[prevN]
				lastSpoken[prevN] = turn - 1
				prevN = turn - prevSpoken - 1
			}
		}
	}

	fmt.Println(prevN)
}
