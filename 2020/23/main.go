package main

import "fmt"

type LinkedList struct {
	n    int
	next *LinkedList
}

func main() {
	//cups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	cups := []int{1, 6, 7, 2, 4, 8, 3, 5, 9}

	l := createLL(cups)
	for move := 1; move <= 100; move++ {
		//printLL(l)
		l = mixUp(l)
		l = l.next
	}

	printLL(find(l, 1))
}

func createLL(nums []int) *LinkedList {
	curr := &LinkedList{n: nums[0]}
	p := curr
	for _, n := range nums[1:] {
		p.next = &LinkedList{n: n}
		p = p.next
	}

	p.next = curr
	return curr
}

func find(p *LinkedList, n int) *LinkedList {
	for p.n != n {
		p = p.next
	}
	return p
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
	p := curr.next

	pickup := []int{}
	for i := 1; i <= 3; i++ {
		pickup = append(pickup, p.n)
		p = p.next
	}

	curr.next = p

	dest := curr.n - 1
	for {
		if has(pickup, dest) {
			dest--
		} else if dest == 0 {
			dest = 9
		} else {
			break
		}
	}

	p = curr
	for {
		if p.n == dest {
			tmp := p.next
			for _, v := range pickup {
				p.next = &LinkedList{n: v}
				p = p.next
			}
			p.next = tmp
			break
		}

		p = p.next
	}

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
