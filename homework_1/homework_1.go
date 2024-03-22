//Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число.
//Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
//При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func create(n int) []int {
	arr := make([]int, n)
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10)
	}
	return arr
}

func findValue(arr []int, size int, value int) {
	result := -1
	for i := 0; i < size; i++ {
		if arr[i] == value {
			fmt.Println("Индекс value:", i)
			fmt.Println("Количество чисел после value:", size-i-1)
			result = i
			break
		}
	}

	if result == -1 {
		fmt.Println("Такого числа нет в массиве, результат:", 0)
	}

}

func main() {
	value := rand.Intn(10)
	fmt.Println("Искомое число:", value)
	arr := create(10)
	size := len(arr)
	fmt.Println(arr)
	findValue(arr, size, value)

}
