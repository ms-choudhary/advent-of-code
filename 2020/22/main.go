package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	player1 := []int{18, 50, 9, 4, 25, 37, 39, 40, 29, 6, 41, 28, 3, 11, 31, 8, 1, 38, 33, 30, 42, 15, 26, 36, 43}
	player2 := []int{32, 44, 19, 47, 12, 48, 14, 2, 13, 10, 35, 45, 34, 7, 5, 17, 46, 21, 24, 49, 16, 22, 20, 27, 23}

	//player1 := []int{9, 2, 6, 3, 1}
	//player2 := []int{5, 8, 4, 7, 10}

	winner, winnerDeck := recursiveCombat(player1, player2)

	score := 0

	for i := 1; i <= len(winnerDeck); i++ {
		score += i * winnerDeck[len(winnerDeck)-i]
	}

	fmt.Println(string(winner) + ": " + strconv.Itoa(score))
}

func toString(deck []int) string {
	strDeck := []string{}
	for _, c := range deck {
		strDeck = append(strDeck, strconv.Itoa(c))
	}

	return strings.Join(strDeck, ":")
}

func key(deckA, deckB []int) string {
	return toString(deckA) + "-" + toString(deckB)
}

func copyArr(src []int) []int {
	res := []int{}
	for _, v := range src {
		res = append(res, v)
	}
	return res
}

func recursiveCombat(deckA, deckB []int) (rune, []int) {

	var lookup = map[string]int{}
	for {
		if _, ok := lookup[key(deckA, deckB)]; ok {
			return 'A', deckA
		} else {
			lookup[key(deckA, deckB)]++
		}

		cardA := deckA[0]
		cardB := deckB[0]

		deckA = deckA[1:]
		deckB = deckB[1:]

		if cardA <= len(deckA) && cardB <= len(deckB) {
			whowon, _ := recursiveCombat(copyArr(deckA[:cardA]), copyArr(deckB[:cardB]))

			switch whowon {
			case 'A':
				deckA = append(deckA, cardA, cardB)
			case 'B':
				deckB = append(deckB, cardB, cardA)
			}
		} else if cardA > cardB {
			deckA = append(deckA, cardA, cardB)
		} else if cardB > cardA {
			deckB = append(deckB, cardB, cardA)
		}

		if len(deckA) == 0 {
			return 'B', deckB
		} else if len(deckB) == 0 {
			return 'A', deckA
		}
	}
}
