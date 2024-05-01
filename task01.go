package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры
	Human (аналог наследования).
*/

// Human structure
type Human struct {
	Name string // Имя
	Age  int    // Возраст
}

// Метод структуры Human для ходьбы
func (p Human) Walk() {
	fmt.Printf("%s is walking \n", p.Name)
}

// Метод структуры Human для бега
func (p Human) Run() {
	fmt.Printf("%s is running \n", p.Name)
}

// Создаем структуру Action на основе встраивания структуры Human
type Action struct {
	Human // Встраиваем структуру Human в структуру Action
}

func task1() {
	fmt.Printf("\n ============ Task 1: ============ \n\n")

	fmt.Printf(" Human: \n")
	person := Human{ //
		Name: "Andrew", // объявляем объект структуры Human
		Age:  22}
	person.Walk()
	person.Run()

	fmt.Printf("\n Action: \n")

	action := Action{person} // объявляем объект структуры Action

	// Вызываем методы Human у объекта структуры Action
	action.Walk()
	action.Run()

	fmt.Printf("\n ================================= \n\n")

}
