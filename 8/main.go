package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// устанавливаем ноль в нужной позиции с помощью
// логического И с нулем
func SetZero(num int64, idx int) int64 {
	return num & ^(1 << (idx - 1))
}

// устанавливаем единицу в нужной позиции с помощью
// логического ИЛИ с единицей
func SetOne(num int64, idx int) int64 {
	return num | (1 << (idx - 1))
}

func main() {
	args := os.Args

	if len(args) != 4 {
		log.Fatal("Should be three arguments: num(int64) idx(1 to 64) bit(0 or 1)")
	}

	num, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	idx, err := strconv.Atoi(args[2]) // 1 to 64
	if err != nil {
		log.Fatal(err.Error())
	}

	bit, err := strconv.Atoi(args[3]) // 0 or 1
	if err != nil {
		log.Fatal(err.Error())
	}

	if bit > 1 || bit < 0 {
		log.Fatal("Bit should be 0 or 1")
	}

	if idx > 64 || idx < 1 {
		log.Fatal("Idx should be from 1 to 64")
	}

	fmt.Printf("before:\t%064b\n", num)

	if bit == 1 {
		num = SetOne(num, idx)
	} else {
		num = SetZero(num, idx)
	}

	fmt.Printf("after:\t%064b\n", num)
}
