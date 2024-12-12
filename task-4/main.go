package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func perebor(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}

func main() {
	var num int
	fmt.Scan(&num)
	var matrix [][]int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < num; i++ {
		readString, _ := reader.ReadString('\n')
		splitString := strings.Split(strings.TrimSpace(readString), " ")
		var raw []int
		for _, num := range splitString {
			atoi, _ := strconv.Atoi(num)
			raw = append(raw, atoi)
		}
		matrix = append(matrix, raw)
	}
	if perebor(matrix) == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
