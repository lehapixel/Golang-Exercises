/* Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m. */

package main

import "fmt"

func main() {

	var a, hours, minutes uint
	fmt.Print("Введите значение угла часовой стрелки(в градусах): ")
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	hours = a / 30
	minutes = 2 * (a % 30)
	fmt.Println("Сейчас", hours, "часов", minutes, "минут.")
}
