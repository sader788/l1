package main

import (
	"fmt"
	"strings"
)

func isUniqeStr(str string) bool {
	lowerStr := strings.ToLower(str)
	existedLet := make(map[rune]struct{})

	for _, let := range lowerStr {
		_, ok := existedLet[let]
		if ok {
			return false
		}
		existedLet[let] = struct{}{}
	}
	return true
}

func main() {
	isUniqe := isUniqeStr("dayr")
	fmt.Println(isUniqe)
}
