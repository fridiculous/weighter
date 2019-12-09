package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLargestWeight(targetWeight float64) float64 {

	weights := [5]float64{45, 25, 10, 5, 2.5}

	var returnWeight float64 = 45.
	for _, key := range weights {
		//fmt.Println(key, returnWeight, targetWeight)
		returnWeight = key
		if key <= targetWeight {
			break
		}
	}

	return returnWeight
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your target weight: ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	target, _ := strconv.ParseFloat(text, 64) //, _ := fmt.Scanln("Please enter a target weight")

	bar := 45.0
	fmt.Println("assuming the bar is", bar)
	remaining := float64(target-float64(bar)) / 2

	final := fmt.Sprintf("%v", bar)

	for remaining >= 2.5 {
		eachWeight := getLargestWeight(remaining)
		fmt.Println("add a", eachWeight, "plate to both sides")

		final = fmt.Sprintf("%s + 2*%v", final, eachWeight)
		remaining = remaining - eachWeight
	}
	total := float64(target) - (remaining * 2)
	fmt.Printf("\n%v = %s\n", total, final)
}
