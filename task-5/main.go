package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	result := strings.Replace(str, "1", "one", -1)
	fmt.Println(result)
}
