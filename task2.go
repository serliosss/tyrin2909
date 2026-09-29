package main

import "fmt"

// 29.09
// ЗАДАЧА №2
// Поменять метсами значение двух переменных
func main1() {
	// 2-1 Через третью переменную
	var a = 10
	var b = 20
	var c = a
	a = b
	b = c
	fmt.Println(a, b)

	// 2-2 Без третьей переменной
	var a1 = 40
	var b1 = 80
	a1 = a1 + b1
	b1 = b1 - a1
	b1 = -b1
	a1 = a1 - b1

	fmt.Println(a1, b1)
}
