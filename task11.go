package main

import "fmt"

/*
   Реализовать пересечение двух неупорядоченных множеств.
*/

func task11() {
	fmt.Printf("\n ============ Task 11: ============ \n\n")

	// создание двух слайсов для примера
	sl1 := []int{1, 2, 3, 4, 5, 6, 10}
	sl2 := []int{7, 8, 5, 4, 109, 10}

	res := make(map[int]bool)
	resInt := make([]int, 0)
	// перебор элементов первого среза и добавление флагов для каждого значения
	for _, value := range sl1 {
		res[value] = true
	}

	// перебор элементов второго среза и проверка, есть ли значение в первом срезе с помощью флагов
	for _, value := range sl2 {
		if res[value] {
			fmt.Printf(" %d  есть в обоих слайсах \n", value)
			resInt = append(resInt, value)
		}
	}
	fmt.Println("\n", resInt)

	fmt.Printf("\n ================================= \n\n")
}
