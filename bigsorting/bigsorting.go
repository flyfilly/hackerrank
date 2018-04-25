package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	inputs := getInput()
	for i := 0; i < len(inputs); i++ {
		fmt.Println(inputs[i])
	}
}

func getInput() []string {
	var input string
	fmt.Scanf("%v", &input)
	var size, _ = strconv.Atoi(strings.TrimSpace(input))
	values := make([]string, size)

	for i := 0; i < size; i++ {
		var val string
		fmt.Scanf("%v", &val)
		values[i] = strings.TrimSpace(val)
	}

	return Sort(values)
}

func Sort(n []string) []string {
	b1 := big.NewInt(0)
	b2 := big.NewInt(0)

	swap := func(i int, j int) {
		tmp := n[i]
		n[i] = n[j]
		n[j] = tmp
	}

	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			out, _ := b1.SetString(n[i], 10)
			in, _ := b2.SetString(n[j], 10)

			if len(n[i]) < len(n[j]) {
				swap(i, j)
			} else if out.Cmp(in) == -1 {
				swap(i, j)
			}
		}
	}

	return n
}
