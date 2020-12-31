package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	exprs := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		exprs = append(exprs, scanner.Text())
	}

	sum := 0
	for _, exp := range exprs {
		sum += eval(exp)
	}

	fmt.Println(sum)
}

func eval(exp string) int {
	mults := []int{}
	acc := 0
	op := '+'
	for i := 0; i < len(exp); {
		if exp[i] == ' ' {
			i++
		} else if unicode.IsDigit(rune(exp[i])) {
			j := i
			for ; j < len(exp); j++ {
				if !unicode.IsDigit(rune(exp[j])) {
					break
				}
			}
			n, err := strconv.Atoi(exp[i:j])
			if err != nil {
				log.Fatalf("not a number %s: %v", exp[i:j], err)
			}
			acc = handle(acc, n, op)
			i = j
		} else if exp[i] == '+' {
			op = rune(exp[i])
			i++
		} else if exp[i] == '*' {
			mults = append(mults, acc)
			acc, op = 0, '+'
			i++
		} else if exp[i] == '(' {
			j, p := i+1, 1
			for ; j < len(exp); j++ {
				if exp[j] == '(' {
					p++
				} else if exp[j] == ')' {
					p--
				}
				if p == 0 {
					break
				}
			}
			//fmt.Println(exp[i+1 : j])

			n := eval(exp[i+1 : j])
			acc = handle(acc, n, op)
			i = j + 1
		}
	}
	mults = append(mults, acc)

	res := 1
	for _, n := range mults {
		res *= n
	}

	return res
}

func handle(a, b int, op rune) int {
	if op == '+' {
		return a + b
	}
	return a * b
}
