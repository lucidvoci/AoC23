package days

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	scanner := util.GetScanner(2)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		power := gamePower(line)
		sum += power
	}
	fmt.Println(sum)

	util.CloseScanner()
}

// Game 1: 1 red, 5 blue, 1 green; 16 blue, 3 red; 6 blue, 5 red; 4 red, 7 blue, 1 green
func gamePower(line string) int {
	var red int
	var green int
	var blue int

	cubes := strings.Split(line, ": ")[1]
	for _, games := range strings.Split(cubes, "; ") {
		for _, game := range strings.Split(games, ", ") {
			num, _ := strconv.Atoi(strings.Split(game, " ")[0])
			color := strings.Split(game, " ")[1]

			if color == "red" && num > red {
				red = num
			} else if color == "green" && num > green {
				green = num
			} else if color == "blue" && num > blue {
				blue = num
			}
		}
	}
	return red * green * blue
}
