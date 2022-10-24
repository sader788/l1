package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	var sb strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		sb.WriteString(words[i])
		sb.WriteByte(' ')
	}
	return sb.String()
}

func main() {
	d := "snow dog sun"

	res := ReverseWords(d)

	fmt.Println(res)
}
