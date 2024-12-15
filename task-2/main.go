package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func triangle(a int, b int, c int) string {
	yes := "YES"
	no := "NO"
	if a+b > c || a+c > b || b+c > a {
		return yes
	}
	return no
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	aStr, _ := reader.ReadString('\n')
	bStr, _ := reader.ReadString('\n')
	cStr, _ := reader.ReadString('\n')

	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	c, _ := strconv.Atoi(cStr)
	result := triangle(a, b, c)
	fmt.Println(result)

}
