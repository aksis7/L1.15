package main

import "fmt"

// Пример функции создания большой строки
func createHugeString(n int) string {
	return string(make([]byte, n))
}

var justString string

func someFunc() {
	// Создаем строку, но сразу берем только первые 100 символов и копируем их в новую строку
	v := createHugeString(1 << 10) // 1024 символа
	justString = string(v[:100])   // Копируем первые 100 символов
}

func main() {
	someFunc()
	fmt.Println(justString) // Выводим результат
}
