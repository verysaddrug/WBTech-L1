package main

import (
	"fmt"
	"sync"
)

/*
   Дана последовательность чисел: 2,4,6,8,10.
   Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func squareSum(sl []int) int {
	sum := 0
	var mu sync.Mutex // Мьютекс для синхронизации доступа к sum
	wg := sync.WaitGroup{}

	// Проходим по всем элементам слайса sl
	for _, v := range sl {
		wg.Add(1) // Увеличиваем счетчик ожидаемых горутин

		// Конкурентно запускаем горутины для работы функции
		go func(n int) {
			defer wg.Done() // Уменьшаем счетчик ожидаемых горутин при завершении
			mu.Lock()       // Блокируем мьютекс перед изменением sum
			sum += n        // Изменяем sum
			mu.Unlock()     // Разблокируем мьютекс после изменения sum
		}(v) // Передаем значение элемента слайса в качестве аргумента для горутины
	}

	wg.Wait()  // Ожидаем завершения всех горутин
	return sum // Возвращаем сумму всех элементов
}

func task3() {
	fmt.Printf("\n ============ Task 3: ============ \n\n")
	sl := []int{2, 4, 6, 8, 10}
	res := squareSum(sl)
	fmt.Println(res) //220
	fmt.Printf("\n ================================= \n\n")
}
