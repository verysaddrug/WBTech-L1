package main

import "fmt"

/*
   Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

// оба метода очень схожи, для множества нам не важно частота повторений слов
// в данной задаче нам важно знать только существует ли ключ
// значением можем быть struct{}{} тк ничего не весит, bool и даже int
// в случае последнего появится доп проверка key >= 1
func methodMapaStruct(str []string) {

	data := make(map[string]struct{}) // мапа с пустой структурой работает как множество

	// проходимся по мапе, все слова-ключи которые будут найдены будут записаны с пустой структурой в мапу data
	for _, word := range str {
		data[word] = struct{}{}
	}
	sl := make([]string, 0, len(data)) //  аллоцируем память для слайса с cap = len(data) - чтобы не аллоцировать лишний раз память

	for word := range data { // записываем все ключи в слайс
		sl = append(sl, word)
	}
	fmt.Println("Метод map и пустая структура:")
	fmt.Printf("%s \n\n", sl)
}

func methodMapaBool(str []string) {

	data := make(map[string]bool)

	for _, word := range str {
		data[word] = true
	}

	sl := make([]string, 0, len(data))

	for word := range data {
		sl = append(sl, word)
	}
	fmt.Println("Метод map и bool:")
	fmt.Printf("%s \n\n", sl)

}

func task12() {
	fmt.Printf("\n ============ Task 12: ============ \n\n")

	str := []string{"cat", "cat", "dog", "cat", "tree"}

	methodMapaStruct(str)

	methodMapaBool(str)

	fmt.Printf("\n ================================= \n\n")
}
