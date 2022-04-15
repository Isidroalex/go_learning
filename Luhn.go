//package luhn

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	id := "4539 3195 0343 6467"
	if len(id) <= 1 {
		fmt.Printf("There is 1 or less chars in string")
	}

	id = strings.Replace(id, " ", "", -1)
	re := regexp.MustCompile("[^0-9]+")
	if len(re.FindAllString(id, -1)) > 0 {
		fmt.Printf("There is chars in the string")
	}

	var intSlice = IntToSlice(id)

	if len(id)%2 == 0 {
		for i := len(id) - 1; i >= 0; i-- {
			for j := 0; j < 2; j++ {
				byteNumber := id[i]
				byteToInt, _ := strconv.Atoi(string(byteNumber))
				fmt.Printf("There is a %d in ther place\n", byteToInt)
			}
		}
	}

	//return true
}

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}
