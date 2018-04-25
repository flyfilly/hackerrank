package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := buildMatrix()
	diff := diagonalDifference(matrix)
	fmt.Println(diff)
}

func buildMatrix() [][]int {
	var input string
	fmt.Scanf("%v", &input)
	var size, _ = strconv.Atoi(strings.TrimSpace(input))
	var matrix = [][]int{}

	for i := 0; i < size; i++ {
		row := bufio.NewReader(os.Stdin)
		line, _ := row.ReadString('\n')
		vals := strings.Split(strings.TrimSpace(line), " ")
		s := make([]int, size)

		for j := 0; j < len(vals); j++ {
			s[j], _ = strconv.Atoi(strings.TrimSpace(vals[j]))
		}

		matrix = append(matrix, s)
	}

	return matrix
}

func diagonalDifference(n [][]int) int {
	var sum1 int
	var sum2 int
	size := len(n)

	for i := 0; i < size; i++ {
		last := (size - 1) - i

		sum1 += n[i][i]
		sum2 += n[i][last]
	}

	diff := (sum1 - sum2)

	if diff < 0 {
		return diff * -1
	}

	return diff
}
