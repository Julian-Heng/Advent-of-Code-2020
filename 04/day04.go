package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	passports, err := readPassports("input")

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(solve(passports, false))
	fmt.Println(solve(passports, true))
}

func solve(passports []passport, extended bool) int {
	result := 0
	for _, p := range passports {
		if p.Validate(extended) {
			result++
		}
	}
	return result
}

func (p passport) Validate(extended bool) bool {
	if p.byr == -1 || p.iyr == -1 || p.eyr == -1 || p.hgt == "" ||
		p.hcl == "" || p.ecl == "" || p.pid == "" {
		return false
	}

	if !extended {
		return true
	}

	if !checkRange(p.byr, 1920, 2002) {
		return false
	}

	if !checkRange(p.iyr, 2010, 2020) {
		return false
	}

	if !checkRange(p.eyr, 2020, 2030) {
		return false
	}

	if strings.HasSuffix(p.hgt, "cm") {
		if !checkRange(atoi(p.hgt[:len(p.hgt)-2]), 150, 193) {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		if !checkRange(atoi(p.hgt[:len(p.hgt)-2]), 59, 76) {
			return false
		}
	} else {
		return false
	}

	if !regexMatch(`^#[0-9A-Za-z]{6}$`, p.hcl) {
		return false
	}

	switch p.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		return false
	}

	if !regexMatch(`^\d{9}$`, p.pid) {
		return false
	}

	return true
}

func readPassports(path string) ([]passport, error) {
	passports := []passport{}
	bfile, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	file := string(bfile)
	delim := getNewline(file)
	split := strings.Split(strings.TrimSpace(file), delim+delim)

	for _, s := range split {
		passports = append(passports, newPassport(s))
	}

	return passports, err
}

func newPassport(str string) passport {
	p := passport{
		byr: -1,
		iyr: -1,
		eyr: -1,
		hgt: "",
		hcl: "",
		ecl: "",
		pid: "",
		cid: "",
	}

	for _, s := range strings.Fields(str) {
		split := strings.SplitN(s, ":", 2)
		k := split[0]
		v := split[1]
		switch k {
		case "byr":
			p.byr = atoi(v)

		case "iyr":
			p.iyr = atoi(v)

		case "eyr":
			p.eyr = atoi(v)

		case "hgt":
			p.hgt = v

		case "hcl":
			p.hcl = v

		case "ecl":
			p.ecl = v

		case "pid":
			p.pid = v

		case "cid":
			p.cid = v
		}
	}

	return p
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}

func checkRange(n int, l int, h int) bool {
	return l <= n && n <= h
}

func regexMatch(p string, s string) bool {
	match, err := regexp.MatchString(p, s)
	if err != nil {
		return false
	}
	return match
}

func getNewline(s string) string {
	for _, v := range s {
		if v == '\r' {
			return "\r\n"
		}
	}
	return "\n"
}
