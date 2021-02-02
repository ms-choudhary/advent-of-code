package main

import "fmt"

func main() {
	player1 := []int{18, 50, 9, 4, 25, 37, 39, 40, 29, 6, 41, 28, 3, 11, 31, 8, 1, 38, 33, 30, 42, 15, 26, 36, 43}
	player2 := []int{32, 44, 19, 47, 12, 48, 14, 2, 13, 10, 35, 45, 34, 7, 5, 17, 46, 21, 24, 49, 16, 22, 20, 27, 23}

	winner := []int{}
	for {
		card1 := player1[0]
		card2 := player2[0]

		player1 = player1[1:]
		player2 = player2[1:]

		if card1 > card2 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}

		if len(player1) == 0 {
			winner = player2
			break
		} else if len(player2) == 0 {
			winner = player1
			break
		}
	}

	score := 0

	for i := 1; i <= len(winner); i++ {
		score += i * winner[len(winner)-i]
	}

	fmt.Println(score)
}
