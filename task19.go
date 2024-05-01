package main

import (
	"fmt"
	"strings"
)

/*
   Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
   Символы могут быть unicode.
*/

func methodBuilder(str string) {
	// создаем билдер для эффективной конкатенации
	var builder strings.Builder
	builder.Grow(len(str))

	//переводим строку из слайсов в руны
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	// создаем строку из инвертированных рун

	fmt.Println("Перевернутая строка methodBuilder:", builder.String())
}

func methodSwap(str string) {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// создаем строку из инвертированных рун
	reversedStr := string(runes)

	fmt.Println("Перевернутая строка methodSwap:", reversedStr)
}

func task19() {
	fmt.Printf("\n ============ Task 19: ============ \n\n")

	// создаем строку и получаем ее значение
	var str string
	fmt.Println("Введите слово:")
	fmt.Scanln(&str)

	methodBuilder(str)

	methodSwap(str)

	fmt.Printf("\n ================================= \n\n")
}
