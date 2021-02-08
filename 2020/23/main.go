package main

import (
	"fmt"
	"time"
)

type LinkedList struct {
	n    int
	next *LinkedList
}

const minN = 1
const maxN = 1000000

var index = map[int]*LinkedList{}

func main() {
	//cups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	cups := []int{1, 6, 7, 2, 4, 8, 3, 5, 9}

	start := time.Now()
	for i := 10; i <= 1000000; i++ {
		cups = append(cups, i)
	}
	fmt.Printf("tt: add list: %v\n", time.Now().Sub(start))

	start = time.Now()
	l := createLL(cups)
	fmt.Printf("tt: create ll: %v\n", time.Now().Sub(start))

	start = time.Now()
	for move := 1; move <= 10000000; move++ {
		//printLL(l)
		//fmt.Println(l.n)
		l = mixUp(l)
		l = l.next
	}
	fmt.Printf("tt: mixup moves: %v\n", time.Now().Sub(start))

	start = time.Now()
	p := find(1)
	fmt.Printf("tt: find key: %v\n", time.Now().Sub(start))

	fmt.Println("ans")

	fmt.Println(p.next.n)
	fmt.Println(p.next.next.n)
	fmt.Println(p.next.n * p.next.next.n)

	//printLL(find(1))
}

func createLL(nums []int) *LinkedList {
	curr := &LinkedList{n: nums[0]}
	index[nums[0]] = curr

	p := curr
	for _, n := range nums[1:] {
		newp := &LinkedList{n: n}

		index[n] = newp
		p.next = newp

		p = p.next
	}

	p.next = curr
	return curr
}

func find(n int) *LinkedList {
	return index[n]
}

func printLL(curr *LinkedList) {

	fmt.Printf("%d", curr.n)

	p := curr.next
	for p != curr {
		fmt.Printf("%d", p.n)
		p = p.next
	}

	fmt.Println()
}

func mixUp(curr *LinkedList) *LinkedList {
	pstart := curr.next
	pend := pstart
	pickup := []int{}

	for i := 1; i <= 2; i++ {
		pickup = append(pickup, pend.n)
		pend = pend.next
	}

	pickup = append(pickup, pend.n)

	curr.next = pend.next

	dest := curr.n - 1
	for {
		if has(pickup, dest) {
			dest--
		} else if dest == 0 {
			dest = maxN
		} else {
			break
		}
	}

	dp := find(dest)
	tmp := dp.next
	dp.next = pstart
	pend.next = tmp

	return curr
}

func has(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}
