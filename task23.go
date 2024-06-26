package main

import "fmt"

func task23() {
	fmt.Printf("\n ============ Task 23: ============ \n\n")

	//создаем слайс
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// запрашиваем индекс для удаления
	var index int
	fmt.Println("Введите индекс элемент для удаления:")
	fmt.Scanln(&index)

	//получаем новый срез, объединяя две части исходного среза:
	//все элементы до указанного индекса и все элементы после указанного индекса.
	//Таким образом, элемент с указанным индексом "вырезается" из среза
	sl = append(sl[:index], sl[index+1:]...)
	fmt.Println("Новый слайс после удаления i-элемента", sl)

	fmt.Printf("\n ================================= \n\n")
}
