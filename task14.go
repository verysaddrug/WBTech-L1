package main

import "fmt"

/*
   Разработать программу, которая в рантайме способна определить тип переменной:
   int, string, bool, channel из переменной типа interface{}.
*/

func task14() {
	fmt.Printf("\n ============ Task 14: ============ \n\n")

	// Объявляем переменные различных типов данных
	var integer int = 1
	var str string = "str"
	var fls bool = false
	var chInt chan int = make(chan int)
	var chStr chan string = make(chan string)

	// Создаем слайс, содержащий переменные различных типов данных
	sl := []interface{}{integer, str, fls, chInt, chStr}

	// Итерируемся по слайсу и выводим тип каждой переменной
	for _, val := range sl {
		switch val.(type) {
		case int:
			fmt.Printf("%d int \n", val)
		case string:
			fmt.Printf("%s string \n", val)
		case bool:
			fmt.Printf("%t bool \n", val)
		case chan int:
			fmt.Printf("chan int \n")
		case chan string:
			fmt.Printf("chan string \n")
		}
	}

	fmt.Printf("\n ================================= \n\n")
}
