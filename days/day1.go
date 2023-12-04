package days

import (
	"advent/util"
	"fmt"
	"math"
	"strings"
)

var stringNums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Day1() {
	scanner := util.GetScanner(1)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		s, e := getNumbers(line)
		sum += concatDigits(s, e)
	}
	fmt.Println(sum)
	util.CloseScanner()
}

func getNumbers(line string) (int, int) {
	s := -1
	e := -1

	si := int(math.MaxInt64)
	ei := 0

	for i := 0; i < len(line); i++ {
		if isNumeric(line[i]) && s == -1 {
			s = byteToInt(line[i])
			si = i
		}
		if isNumeric(line[len(line)-i-1]) && e == -1 {
			e = byteToInt(line[len(line)-i-1])
			ei = len(line) - i - 1
		}
		if s != -1 && e != -1 {
			break
		}
	}
	for j, c := range stringNums {
		k := strings.Index(line, c)
		l := strings.LastIndex(line, c)

		if k != -1 && k < si {
			si = k
			s = j + 1
		}
		if l != -1 && l > ei {
			ei = l
			e = j + 1
		}
	}
	return s, e
}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}

func byteToInt(b byte) int {
	return int(b - '0')
}

func concatDigits(d1 int, d2 int) int {
	return d1*10 + d2
}
