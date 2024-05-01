package main

import "fmt"

/*
   Реализовать бинарный поиск встроенными методами языка.

   Для корректной работы слайс должен быть отсортирован!
*/

func binarySearch(sl []int, target int) int {
	left, right := 0, len(sl)-1

	for left <= right { // продолжаем, пока левая граница не превысит правую
		mid := left + (right-left)/2 // вычисляем индекс середины
		if target == mid {           // если цель равна середине, возвращаем индекс середины
			return mid
		} else if mid < target { // если цель больше середины, сдвигаем левую границу
			left = mid + 1
		} else { // если цель меньше середины, сдвигаем правую границу
			right = mid - 1
		}
	}
	return -1 // если цель не найдена, возвращаем -1
}

func task17() {
	fmt.Printf("\n ============ Task 17: ============ \n\n")

	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var target int
	fmt.Println("Введите элемент для поиска:")
	fmt.Scanln(&target)

	res := binarySearch(sl, target)
	if res == -1 {
		fmt.Printf("Элемент %d не найден..", target)
	} else {
		fmt.Printf("Элемент %d найден под индеком %d", target, res-1)
	}

	fmt.Printf("\n ================================= \n\n")
}
