package main

import (
	"fmt"
	"math/big"
)

/*
   Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
   значение которых > 2^20.

   Возможно опечатка, тк 2^20 спокойно влазит в int64
*/

func calculateMistake() {
	fmt.Println("Method bigInt:")
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений числовых переменных
	a.SetString("200000000000000000000", 10)
	b.SetString("300000000000000000000", 10)

	// Сложение a и b
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма a и b:", sum)

	// Вычитание b из a
	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность a и b:", diff)

	// Умножение a на b
	prod := new(big.Int).Mul(a, b)
	fmt.Println("Произведение a и b:", prod)

	// Проверка деления на ноль
	if b.Sign() == 0 {
		fmt.Println("Деление на ноль!")
		return
	}

	// Деление a на b с округлением вниз
	quotient := new(big.Int).Div(a, b)
	fmt.Printf("Частное a и b: %s \n\n", quotient)
}

func calculateWithMistake() {
	// Задаем значения переменных a и b
	fmt.Println("Method int64:")
	a := int64(2000000) // 2^20
	b := int64(3000000) // 3 * 2^20

	// Сложение a и b
	sum := a + b
	fmt.Println("Сумма a и b:", sum)

	// Вычитание b из a
	diff := a - b
	fmt.Println("Разность a и b:", diff)

	// Умножение a на b
	prod := a * b
	fmt.Println("Произведение a и b:", prod)

	// Проверка деления на ноль
	if b == 0 {
		fmt.Println("Деление на ноль!")
	} else {
		// Деление a на b
		quotient := a / b
		fmt.Println("Частное a и b:", quotient)
	}
}

func task22() {
	fmt.Printf("\n ============ Task 22: ============ \n\n")

	calculateMistake() // если в тз ошибка и там 2e20

	calculateWithMistake() // если в тз нет ошибки

	fmt.Printf("\n ================================= \n\n")
}
