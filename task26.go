package main

import "fmt"

/*
   Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
   Функция проверки должна быть регистронезависимой.

    Например:
    abcd — true
    abCdefAaf — false
	aabcd — false

*/

func task26() {
	fmt.Printf("\n ============ Task 26: ============ \n\n")

	//получаем строку
	var str string
	fmt.Println("Введите строку:")
	fmt.Scanln(&str)

	// создаем мапу для записи каждого символа
	data := make(map[rune]int)

	// проходимся по строке, если ключ-символ встречается добавляем +1 value
	for _, c := range str {
		data[c] += 1

		// если value  по какому либо ключу >1, то слово уже не уникально, выводим false и return'имя
		if data[c] > 1 {
			fmt.Println("False")
			fmt.Printf("\n ================================= \n\n")
			return
		}

	}
	// если не встретилось ключа со значением >1, то слово уникально и выводим true
	fmt.Println("True")

	fmt.Printf("\n ================================= \n\n")
}
