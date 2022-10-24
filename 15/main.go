package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	sb := strings.Builder{}

	for i := 0; i < size; i++ {
		sb.WriteRune('ꡄ')
	}

	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	//Срез строки делаем из байтов, но размер руны может быть больше одного байта
	//Следовательно мы можем получить меньшее количество символов, чем есть
	//Либо один из символов будет обрезан
	justString := v[:100]
	//Чтобы сделать срез правильно, нужно привести строку к слайсу рун
	runes := []rune(v)
	str := string(runes[:100])

	fmt.Println(justString)
	fmt.Println(str)
}

func main() {
	someFunc()
}
