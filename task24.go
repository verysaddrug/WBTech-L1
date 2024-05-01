package main

import (
	"fmt"
	"math"
)

/*
   Разработать программу нахождения расстояния между двумя точками,
   которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

// структура координат точки
type Point struct {
	X, Y int
}

// конструктор новой точки
func newPoint() Point {
	return Point{0, 0}
}

// функция нахождения дистанции двух точек используя теорему Пифагора
func distance(p1, p2 *Point) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}

func task24() {
	fmt.Printf("\n ============ Task 24: ============ \n\n")

	// создаем две точки используя конструктор
	p1 := newPoint()
	p2 := newPoint()

	p2.X = 7
	p2.Y = 5

	fmt.Printf("Дистанция: %f\n", distance(&p1, &p2))

	fmt.Printf("\n ================================= \n\n")
}
