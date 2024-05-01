package main

import "fmt"

//   Поменять местами два числа без создания временной переменной.

func task13() {
	fmt.Printf("\n ============ Task 13: ============ \n\n")

	// инициализация
	var a, b = 1, 2
	fmt.Printf("a = %d, b = %d\n", a, b)

	// swap
	a, b = a, b
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Printf("\n ================================= \n\n")
}
