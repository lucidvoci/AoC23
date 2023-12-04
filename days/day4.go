package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() {
	filePath := "./inputs/input4.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res [206]int
	for i := range res {
		res[i] = 1
	}

	var gameNum int
	for scanner.Scan() {
		line := scanner.Text()
		copies(line, gameNum, &res)
		gameNum++
	}

	var sum int
	for i := range res {
		sum += res[i]
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func copies(line string, gameNum int, res *[206]int) {
	win := make(map[string]bool)
	game := strings.Split(line, ":")
	nums := strings.Split(game[1], "|")
	for _, winNum := range strings.Split(nums[0], " ") {
		cleanNum := strings.TrimSpace(winNum)
		_, err := strconv.Atoi(cleanNum)
		if err != nil {
			continue
		}
		win[cleanNum] = true
	}
	var match int
	for _, myNum := range strings.Split(nums[1], " ") {
		cleanNum := strings.TrimSpace(myNum)
		if win[cleanNum] {
			match++
		}
	}
	for i := gameNum + 1; i < gameNum+match+1 && i < 206; i++ {
		res[i] += res[gameNum]
	}
}
