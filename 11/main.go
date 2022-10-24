package main

import "fmt"

func Intersection[T comparable](a1 []T, a2 []T) []T {
	intersected := make(map[T]struct{})
	//Записываем все значения первого слайса в хэш
	for _, el := range a1 {
		intersected[el] = struct{}{}
	}

	var result []T
	//Если элемент второго слайса уже есть в хэше
	//Значит пересечение, добавляем в результат
	for _, el := range a2 {
		if _, ok := intersected[el]; ok == true {
			result = append(result, el)
		}
	}

	return result
}

func main() {
	a1 := []int{5, 3, 2, 1}
	a2 := []int{2, 4, 7, 1, 6, 5, 10, 11}

	res := Intersection(a1, a2)

	fmt.Println(res)

}
