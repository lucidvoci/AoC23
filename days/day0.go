package days

import (
	"advent/util"
	"fmt"
)

func Day0() {
	scanner := util.GetScanner(0)

	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		sum += 0
	}
	fmt.Println(sum)

	util.CloseScanner()
}
