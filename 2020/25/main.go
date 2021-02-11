package main

import "fmt"

func main() {
	//cardPubKey := 5764801
	//doorPubKey := 17807724
	cardPubKey := 13233401
	doorPubKey := 6552760

	cardLoopSize := loopSize(cardPubKey)
	fmt.Println(cardLoopSize)

	doorLoopSize := loopSize(doorPubKey)
	fmt.Println(doorLoopSize)

	fmt.Println(loop(cardPubKey, doorLoopSize))
}

func loopSize(pubkey int) int {
	val := 1

	loopSize := 0
	for {
		val *= 7
		val = val % 20201227
		loopSize++

		if val == pubkey {
			return loopSize
		}
	}

	return loopSize
}

func loop(subjectNumber, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= subjectNumber
		val = (val % 20201227)
	}
	return val
}
