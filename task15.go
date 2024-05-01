package main

import (
	"fmt"
	"strings"
)

/*
   К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
   Приведите корректный пример реализации.

   var justString string
   func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
   }

   func main() {
    someFunc()
   }

Данный фрагмент кода содержит проблему, связанную с неэффективным использованием памяти.
Функция someFunc создаёт большую строку v с помощью функции createHugeString,
но затем из этой строки в глобальную переменную justString сохраняется только первые 100 символов.
Однако из-за того, как устроено управление памятью в Go, вся строка v остаётся в памяти,
так как срез строки сохраняет ссылку на исходный массив байтов.
Таким образом, несмотря на то, что нам нужны только 100 символов,
в памяти сохраняется вся большая строка, что ведет к избыточному расходу памяти.
*/

func someFunc() string {
	v := createHugeString(1 << 10)
	justString := string([]byte(v[:100])) // делам локальную переменную с 100 первыми символами  из v
	return justString
}

func createHugeString(size int) string {
	// создание билдера строк с заданным размером
	var builder strings.Builder
	builder.Grow(size)

	// заполнение билдером строк символом 'a'
	for i := 0; i < size; i++ {
		builder.WriteRune('a')
	}

	// возврат строки билдера
	return builder.String()
}

func task15() {
	fmt.Printf("\n ============ Task 15: ============ \n\n")

	fmt.Printf("justString: %v\n", someFunc())

	fmt.Printf("\n ================================= \n\n")
}
