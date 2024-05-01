package main

import (
	"context"
	"fmt"
	"sync"
)

/*
   Реализовать конкурентную запись данных в map.
*/

func setMap(data map[int]int, ch chan int, ctx context.Context, mu *sync.Mutex, wg *sync.WaitGroup) {
	for {
		select {
		case dt := <-ch:
			// блокировка мапы с использованием мьютекса
			mu.Lock()

			// сохранение данных
			data[dt] = dt
			// разблокировка мапы с использованием мьютекса
			mu.Unlock()
			fmt.Printf("Значение %v сохранено\n", dt)
		case <-ctx.Done():
			fmt.Printf("Горутина завершилась\n")
			wg.Done()
			return
		}
	}
}

func task7() {
	fmt.Printf("\n ============ Task 7: ============ \n\n")

	// создание мапы
	data := make(map[int]int)

	// создание канала для отправки данных
	ch := make(chan int)
	defer close(ch)

	// создание мьютекса для предотвращения гонок данных
	mu := sync.Mutex{}

	// создание контекста для остановки горутин
	ctx, cancel := context.WithCancel(context.Background())

	// создание группы ожидания для ожидания завершения всех рабочих
	var wg sync.WaitGroup
	wg.Add(2)

	// создание двух горутин
	for i := 0; i < 2; i++ {
		go setMap(data, ch, ctx, &mu, &wg)
	}

	// отправка данных в канал
	for i := 0; i < 15; i++ {
		ch <- i
	}

	// отправка сигнала о завершении горутин
	cancel()

	// ожидание завершения всех горутин
	wg.Wait()

	fmt.Printf("\n ================================= \n\n")
}
