package main

import (
	"fmt"
	"sort"
)

func TemperatureGroup(temps []float64) map[int][]float64 {
	sort.Float64s(temps)

	var result map[int][]float64 = make(map[int][]float64)
	//Откдиываем последнюю целую цифру числа и умножаем на 10
	//Полученный результат будет идентификатором группы
	for _, t := range temps {
		group := (int(t) / 10) * 10
		result[group] = append(result[group], t)
	}
	return result
}

func main() {
	temperatures := []float64{-25.4, -15, -16.6, -19, -20, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := TemperatureGroup(temperatures)
	fmt.Println(groups)
}
