package main

import "fmt"

func NewSet(strings []string) []string {
	existed := make(map[string]struct{})
	// Аналогично с прошлой задачей, загоняем все элементы слайса в хэш
	// Так как ключ уникальное значение, мы исключаем повторение
	for _, str := range strings {
		existed[str] = struct{}{}
	}

	result := make([]string, 0, len(existed))
	// Добавляем все значение map'a в результат
	for key, _ := range existed {
		result = append(result, key)
	}
	return result
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet(strings)
	fmt.Println(set)
}
