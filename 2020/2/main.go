package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	validPassCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), ":")

		if validPassword(newPolicy(v[0]), v[1][1:]) {
			validPassCount++
		}
	}

	fmt.Println(validPassCount)
}

type PasswordPolicy struct {
	min, max int
	musthave string
}

func newPolicy(s string) PasswordPolicy {
	v := strings.Split(s, " ")
	nos := strings.Split(v[0], "-")

	min, err := strconv.Atoi(nos[0])
	if err != nil {
		log.Fatal(err)
	}

	max, err := strconv.Atoi(nos[1])
	if err != nil {
		log.Fatal(err)
	}

	return PasswordPolicy{min: min, max: max, musthave: v[1]}
}

func validPassword(policy PasswordPolicy, password string) bool {
	return (string(password[policy.min-1]) == policy.musthave && string(password[policy.max-1]) != policy.musthave) || (string(password[policy.min-1]) != policy.musthave && string(password[policy.max-1]) == policy.musthave)
}
