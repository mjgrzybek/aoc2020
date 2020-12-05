package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func loadData(reader io.Reader) [][]string {
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.SplitFunc(SplitAt("\n\n")))

	data := make([][]string, 0)
	for scanner.Scan() {
		data = append(data, strings.Fields(scanner.Text()))
	}

	return data
}

func SplitAt(substring string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte(substring)
	searchLen := len(searchBytes)

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		if atEOF {
			return dataLen, data, nil
		}

		return 0, nil, nil
	}
}

func Solve1() int {
	f, _ := os.Open("day4/input.txt")
	data := loadData(bufio.NewReader(f))

	obligatoryFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	sort.Strings(obligatoryFields)

	valid := 0

	for _, row := range data {
		fields := make([]string, 0)
		for _, kv := range row {
			s := strings.SplitN(kv, ":", 2)
			if s[0] != "cid" {
				fields = append(fields, s[0])
			}
		}
		sort.Strings(fields)

		if reflect.DeepEqual(obligatoryFields, fields) {
			valid++
		}
	}

	return valid
}

type Passport struct {
	Byr, Iyr, Eyr           int
	Hgt, Hcl, Ecl, Cid, Pid string
}

func (p Passport) isValid() bool {
	return p.isValidByr() && p.isValidIyr() && p.isValidEyr() && p.isValidHgt() && p.isValidPid() && p.isValidHcl() && p.isValidEcl() && p.isValidCid()
}

func (p Passport) isValidByr() bool {
	return p.Byr >= 1920 && p.Byr <= 2002
}

func (p Passport) isValidIyr() bool {
	return p.Iyr >= 2010 && p.Iyr <= 2020
}

func (p Passport) isValidEyr() bool {
	return p.Eyr >= 2020 && p.Eyr <= 2030
}

func (p Passport) isValidHgt() bool {
	if len(p.Hgt) == 0 {
		return false
	}

	value, _ := strconv.Atoi(p.Hgt[:len(p.Hgt)-2])
	unit := p.Hgt[len(p.Hgt)-2:]

	return (unit == "cm" && value >= 150 && value <= 193) || (unit == "in" && value >= 59 && value <= 76)
}

func (p Passport) isValidPid() bool {
	if len(p.Pid) != 9 {
		return false
	}

	var sink int
	_, err := fmt.Sscan(p.Pid, &sink)
	return err == nil
}

func (p Passport) isValidHcl() bool {
	match, _ := regexp.MatchString("#[0-9a-f]{6}", p.Hcl)
	return match
}

func (p Passport) isValidEcl() bool {
	for _, s := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if p.Ecl == s {
			return true
		}
	}
	return false
}

func (p Passport) isValidCid() bool {
	return true
}

func Solve2() int {
	f, _ := os.Open("day4/input.txt")
	data := loadData(bufio.NewReader(f))

	valid := 0
	for _, passdata := range data {
		if NewPassport(passdata).isValid() {
			valid++
		}
	}
	return valid
}

func NewPassport(passdata []string) Passport {
	p := Passport{}
	for _, x := range passdata {
		kv := strings.SplitN(x, ":", 2)
		key := kv[0]
		value := kv[1]

		switch key {
		case "byr":
			atoi, _ := strconv.Atoi(value)
			p.Byr = atoi
		case "iyr":
			atoi, _ := strconv.Atoi(value)
			p.Iyr = atoi
		case "eyr":
			atoi, _ := strconv.Atoi(value)
			p.Eyr = atoi
		case "hgt":
			p.Hgt = value
		case "pid":
			p.Pid = value
		case "hcl":
			p.Hcl = value
		case "ecl":
			p.Ecl = value
		case "cid":
			p.Cid = value
		}
	}

	return p
}
