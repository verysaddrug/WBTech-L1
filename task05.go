package main

import (
	"context"
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func sender(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func receiver(ctx context.Context, ch <-chan int) {
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				// Канал закрыт, завершаем работу
				return
			}
			fmt.Printf("Полученное значение: %d\n", value)
		case <-ctx.Done():
			// Получен сигнал завершения, завершаем работу
			return
		}
	}
}

func task5() {
	fmt.Printf("\n ============ Task 5: ============ \n\n")

	ch := make(chan int)                                    // Канал для передачи данных
	ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст и функцию для отмены
	timeout := 5 * time.Second                              // Время выполнения
	timeoutChan := time.After(timeout)                      // Таймер с кейсом времени N секунд

	// Запускаем горутину для отправки значений в канал
	go sender(ch)

	// Запускаем горутину для чтения из канала
	go receiver(ctx, ch)

	// Ожидаем сигнала завершения или истечения времени
	select {
	case <-timeoutChan:
		// Истекло время выполнения, отменяем контекст
		fmt.Println("Время выполнения истекло.")
		cancel()
	}

	// Ждем некоторое время, чтобы дать горутинам sender и receiver возможность завершиться
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\n ================================= \n\n")
}
