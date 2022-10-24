package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		log.Fatal("Should be three arguments: num1 num2 opetation_name")
	}

	num1, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	num2, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	// используем пакет big для работы с большими числами
	var res *big.Int = big.NewInt(0)

	switch args[3] {
	case "mul":
		res = res.Mul(big.NewInt(num1), big.NewInt(num2))
	case "div":
		res = res.Quo(big.NewInt(num1), big.NewInt(num2))
	case "add":
		res = res.Add(big.NewInt(num1), big.NewInt(num2))
	case "sub":
		res = res.Sub(big.NewInt(num1), big.NewInt(num2))
	}

	fmt.Println("Result is: ", res.String())
}
