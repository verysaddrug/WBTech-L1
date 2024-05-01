package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

func worker(id int, dataChan <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				// Канал dataChan закрыт, завершаем работу
				return
			}
			time.Sleep(100 * time.Millisecond) // Задержка в 1.5 секунды
			fmt.Printf("Воркер %d: Полученные данные %d\n", id, data)
		case <-ctx.Done():
			// Получен сигнал завершения, завершаем работу
			return
		}
	}
}

func task4() {
	fmt.Printf("\n ============ Task 4: ============ \n\n")

	var N int
	fmt.Println("Введите количество воркеров:")
	fmt.Scanln(&N)

	dataChan := make(chan int) // Канал для передачи данных от главного потока к воркерам

	var wg sync.WaitGroup
	wg.Add(N)

	// Создаем контекст для отмены выполнения воркеров
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запускаем N воркеров
	for i := 1; i <= N; i++ {
		go worker(i, dataChan, &wg, ctx)
	}

	// Запись данных в канал из главного потока
	go func() {
		for i := 1; ; i++ {
			dataChan <- i
		}
	}()

	// Ожидаем сигнала завершения (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-sigChan

	// Отменяем контекст, чтобы завершить выполнение воркеров
	cancel()

	// Дожидаемся завершения работы всех воркеров
	wg.Wait()

	// Закрываем канал, чтобы завершить операции чтения из dataChan
	close(dataChan)

	fmt.Printf("\n ================================= \n\n")
}
