package main

import (
	"fmt"
	"sync"
)

/*
   Написать программу, которая конкурентно рассчитает значение квадратов
   чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func squareVal(nums []int) {
	// WaitGroup для синхронизации
	wg := sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1) // Увеличиваем счетчик ожидаемых горутин

		// конкурентно запускаем горутины для работы функции
		go func(n int) { // пробрасываем n для локального создания копии n для каждой горутины и предотвращения гонок
			defer wg.Done() // Уменьшаем счетчик ожидаемых горутин при завершении

			fmt.Println(n * n) // Выводим результат
		}(v)
	}

	wg.Wait() // Ожидаем завершения всех горутин
}

func task2() {
	fmt.Printf("\n ============ Task 2: ============ \n\n")

	sl := []int{2, 4, 6, 8, 10}
	squareVal(sl)

	fmt.Printf("\n ================================= \n\n")
}
