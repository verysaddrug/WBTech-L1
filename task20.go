package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
   Разработать программу, которая переворачивает слова в строке.
   Пример: «snow dog sun — sun dog snow».
*/

func task20() {
	fmt.Printf("\n ============ Task 20: ============ \n\n")

	var str string
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n') // Считываем ввод пользователя до символа новой строки
	if err != nil {
		fmt.Printf("Ошибка: %v \n", err) // Выводим ошибку, если она есть
	}

	var builder strings.Builder // Создаем билдер строк для эффективной манипуляции строками
	builder.Grow(len(str))      // Предварительно выделяем память в билдере строк для избежания перераспределений

	str = strings.TrimSpace(str) // Удаляем начальные и конечные пробелы из введенной строки
	words := strings.Fields(str) // Разбиваем строку на слова

	// Инвертируем порядок слов и формируем инвертированную строку
	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i]) // Добавляем каждое слово в построитель строк
		builder.WriteRune(' ')        // Добавляем пробел после каждого слова
	}

	fmt.Printf("Перевернутая строка: %s \n", builder.String()) // Выводим инвертированную строку

	fmt.Printf("\n ================================= \n\n")
}
