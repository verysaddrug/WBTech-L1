package main

import (
	"fmt"
	"time"
)

/*
   Реализовать собственную функцию sleep.
*/

func methodSleepTimeNow(sec int) {
	// getting the current time
	start := time.Now()

	for {
		if time.Since(start).Seconds() > float64(sec) {
			fmt.Printf("ВАДИМ, БУДИЛЬНИК ЗВОНИТ, АЛЕ ВАДИМ, ВАДИМ, ВАДИМ!!! АЛЁ!!! \n\n")
			break
		}
	}
}

func methodSleepTimer(sec int) {
	timer := time.NewTimer(time.Duration(sec) * time.Second) //
	<-timer.C
}

func methodSleepSelect(sec int) {
	select {
	case <-time.After(time.Duration(sec) * time.Second):
		return
	}
}

func task25() {
	fmt.Printf("\n ============ Task 25: ============ \n\n")

	var sec int
	fmt.Println(" Введите количество секунд для сна:")
	fmt.Scanln(&sec)

	methodSleepTimeNow(sec)
	fmt.Println("Сон после TimeNow \n")

	methodSleepTimer(sec)
	fmt.Println("Сон после Timer \n")

	methodSleepSelect(sec) // схожая концепция с timer
	fmt.Println("Сон после Select")

	fmt.Printf("\n ================================= \n\n")
}
