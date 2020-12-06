package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var validECL = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(f)
	}

	scanner := bufio.NewScanner(f)

	passports := []Passport{}
	passStr := ""
	for scanner.Scan() {
		line := scanner.Text()

		if "" == line {
			p, err := newPassport(passStr)
			if err == nil {
				passports = append(passports, p)
			}
			passStr = ""
		} else if "" == passStr {
			passStr = line
		} else {
			passStr += " " + line
		}
	}

	p, err := newPassport(passStr)
	if err == nil {
		passports = append(passports, p)
	}

	fmt.Println(len(passports))
}

type Passport struct {
	byr, iyr, eyr int
	hgt           int
	hcl, ecl      string
	pid, cid      string
}

func numericBtn(val string, low, high int) (int, error) {
	n, err := strconv.Atoi(val)
	if err != nil {
		return -1, err
	}
	if n < low || n > high {
		return -1, errors.New("invalid")
	}
	return n, nil
}

func isNumeric(c rune) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

func toString(p Passport) string {
	return fmt.Sprintf("byr: %v, iyr: %v, eyr: %v, hgt: %v, hcl: %v, ecl: %v, pid: %v", p.byr, p.iyr, p.eyr, p.hgt, p.hcl, p.ecl, p.pid)
}

func newPassport(s string) (Passport, error) {
	p := Passport{byr: -1, iyr: -1, eyr: -1, hgt: -1, hcl: "", ecl: "", pid: ""}
	fields := strings.Split(s, " ")

	for _, f := range fields {

		key, val := strings.Split(f, ":")[0], strings.Split(f, ":")[1]

		switch key {
		case "byr":
			yr, err := numericBtn(val, 1920, 2002)
			if err != nil {
				return Passport{}, err
			}

			p.byr = yr

		case "iyr":
			yr, err := numericBtn(val, 2010, 2020)
			if err != nil {
				return Passport{}, err
			}

			p.iyr = yr

		case "eyr":
			yr, err := numericBtn(val, 2020, 2030)
			if err != nil {
				return Passport{}, err
			}

			p.eyr = yr

		case "hgt":
			switch val[len(val)-2:] {
			case "cm":
				hgt, err := numericBtn(val[:len(val)-2], 150, 193)
				if err != nil {
					return Passport{}, err
				}
				p.hgt = hgt

			case "in":
				hgt, err := numericBtn(val[:len(val)-2], 59, 76)
				if err != nil {
					return Passport{}, err
				}

				p.hgt = int(float64(hgt) * 2.54)
			default:
				return Passport{}, errors.New("invalid hgt")
			}

		case "hcl":
			if val[0] != '#' {
				return Passport{}, errors.New("invalid hcl")
			}
			if len(val[1:]) != 6 {
				return Passport{}, errors.New("invalid hcl")
			}

			for _, c := range val[1:] {
				if !isNumeric(c) && c != 'a' && c != 'b' && c != 'c' && c != 'd' && c != 'e' && c != 'f' {
					return Passport{}, errors.New("invalid hcl")
				}
			}

			p.hcl = val

		case "ecl":
			found := false
			for _, v := range validECL {
				if val == v {
					found = true
				}
			}

			if !found {
				return Passport{}, errors.New("invalid ecl")
			}

			p.ecl = val

		case "pid":
			if len(val) != 9 {
				return Passport{}, errors.New("invalid pid")
			}
			for _, c := range val {
				if !isNumeric(c) {
					return Passport{}, errors.New("invalid pid")
				}
			}

			p.pid = val

		case "cid":
			p.cid = val
		}
	}

	if p.byr == -1 || p.iyr == -1 || p.eyr == -1 || p.hgt == -1 || p.hcl == "" || p.ecl == "" || p.pid == "" {
		return Passport{}, errors.New("missing mandatory fields")
	}

	return p, nil
}
