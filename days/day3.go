package days

import (
	"bufio"
	"fmt"
	"os"
)

const s int = 140

func Day3() {
	filePath := "./inputs/input3.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	var schema [s][s]byte
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		createSchema(&schema, line, lineNum)
		lineNum++
	}
	sum = numSum(&schema)
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

type Spot struct {
	Schema *[s][s]byte
	I      int
	J      int
}

func numSum(schema *[s][s]byte) int {
	var sum int
	gearMap := make(map[int]int)
	curNum := -1
	curGear := -1
	for i := 0; i < s; i++ {
		if curNum != -1 && curGear != -1 {
			val, ok := gearMap[curGear]
			if ok {
				sum += curNum * val
			} else {
				gearMap[curGear] = curNum
			}
		}
		curNum = -1
		curGear = -1

		for j := 0; j < s; j++ {
			c := schema[i][j]
			if isNumeric(c) && curGear == -1 {
				curGear = checkIfIsWithGear(Spot{schema, i, j})
			}
			if isNumeric(c) && curNum == -1 {
				curNum = int(c - '0')
			} else if isNumeric(c) && curNum != -1 {
				curNum = curNum*10 + int(c-'0')
			} else {
				if curNum != -1 && curGear != -1 {
					val, ok := gearMap[curGear]
					if ok {
						sum += curNum * val
					} else {
						gearMap[curGear] = curNum
					}

				}
				curNum = -1
				curGear = -1
			}
		}
	}
	return sum
}

func checkIfIsWithGear(spot Spot) int {
	nextSpots := []Spot{
		{spot.Schema, spot.I - 1, spot.J - 1},
		{spot.Schema, spot.I - 1, spot.J},
		{spot.Schema, spot.I, spot.J - 1},
		{spot.Schema, spot.I + 1, spot.J + 1},
		{spot.Schema, spot.I + 1, spot.J},
		{spot.Schema, spot.I, spot.J + 1},
		{spot.Schema, spot.I - 1, spot.J + 1},
		{spot.Schema, spot.I + 1, spot.J - 1},
	}

	for _, nextSpot := range nextSpots {
		curGear := checkGear(nextSpot)
		if curGear != -1 {
			return curGear
		}
	}
	return -1
}

func checkGear(spot Spot) int {
	if spot.I < 0 || spot.J < 0 || spot.I >= s || spot.J >= s {
		return -1
	} else if spot.Schema[spot.I][spot.J] == '*' {
		return spot.I*1000 + spot.J
	}
	return -1
}

func createSchema(schema *[s][s]byte, line string, lineNum int) {
	for i, r := range line {
		schema[lineNum][i] = byte(r)
	}
}
