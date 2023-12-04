package days

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

const rs int = 206

func Day4() {
	scanner := util.GetScanner(4)

	var res [rs]int
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

	util.CloseScanner()
}

func copies(line string, gameNum int, res *[rs]int) {
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
