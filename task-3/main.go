package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func slices(numbers []int) []int {
	slice := make([]int, 0, len(numbers))
	slice = append(slice, numbers[len(numbers)-1])
	slice = append(slice, numbers[:len(numbers)-1]...)
	return slice
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	reader := bufio.NewReader(os.Stdin)
	readStr, _ := reader.ReadString('\n')
	result := strings.Split(readStr, " ")
	numbers := make([]int, 0, num)
	for _, ch := range result {
		atoi, _ := strconv.Atoi(strings.TrimSpace(ch))
		numbers = append(numbers, atoi)
	}
	fmt.Println(slices(numbers))
}
