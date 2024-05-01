package main

import "fmt"

func main() {
	var taskNum string
	fmt.Println("Введите номер задания 1-26 ('exit или  -1' для выхода):")
	fmt.Scan(&taskNum)

	for taskNum != "exit" && taskNum != "-1" {
		switch taskNum {
		case "1":
			task1()
		case "2":
			task2()
		case "3":
			task3()
		case "4":
			task4()
		case "5":
			task5()
		case "6":
			task6()
		case "7":
			task7()
		case "8":
			task8()
		case "9":
			task9()
		case "10":
			task10()
		case "11":
			task11()
		case "12":
			task12()
		case "13":
			task13()
		case "14":
			task14()
		case "15":
			task15()
		case "16":
			task16()
		case "17":
			task17()
		case "18":
			task18()
		case "19":
			task19()
		case "20":
			task20()
		case "21":
			task21()
		case "22":
			task22()
		case "23":
			task23()
		case "24":
			task24()
		case "25":
			task25()
		case "26":
			task26()
		default:
			fmt.Println("Неверный номер задания или команда")
		}
		fmt.Println("Введите номер задания ('exit или  -1' для выхода):")
		fmt.Scan(&taskNum)
	}
}
