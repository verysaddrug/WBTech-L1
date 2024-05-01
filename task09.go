package main

import (
	"context"
	"fmt"
	"sync"
)

/*
   Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
   во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func double(valueChannel chan int, doubleChannel chan int, ctx context.Context) {
	for {
		select {
		case value := <-valueChannel:
			doubleChannel <- value * 2
		case <-ctx.Done():
			return
		}
	}
}

func printDouble(doubleChannel chan int, wg *sync.WaitGroup, ctx context.Context) {
	for {
		select {
		case value := <-doubleChannel:
			fmt.Printf("%v is a double \n", value)
			wg.Done()

		case <-ctx.Done():
			return
		}
	}
}

func task9() {
	fmt.Printf("\n ============ Task 9: ============ \n\n")

	var wg sync.WaitGroup

	valueChannel := make(chan int)
	doubleChannel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go double(valueChannel, doubleChannel, ctx)
	go printDouble(doubleChannel, &wg, ctx)

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		valueChannel <- i
	}

	wg.Wait()

	cancel()

	fmt.Printf("\n ================================= \n\n")
}
