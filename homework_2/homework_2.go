//Заполните упорядоченный массив из 12 элементов и введите число.
//Необходимо реализовать поиск первого вхождения заданного числа в массив. Сложность алгоритма должна быть минимальная.
//При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 12

func create() [size]int {
	var arr [size]int
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = 10*i + rand.Intn(3)
	}
	fmt.Println(arr)
	return arr
}

//Сначала написала это, а потом поняла что нужно первое вхождение найти.. Это как-то сложнее выглядит для первого вхождения

/*
func findValue(arr [size]int, value int) {
	min := 0
	max := size - 1
	result := -1
	for max >= min {
		middle := (max + min) / 2
		if arr[middle] > value {
			max = middle - 1
		} else if arr[middle] < value {
			min = middle + 1
		} else {
			result = middle
			break
		}
	}
	if result == -1 {
		fmt.Println("Искомого значения нет в массиве")
	} else {
		fmt.Println("Индекс искомого значения:", result)
	}

}
*/

func findValue(arr [size]int, value int) {
	k := -1
	for i := 0; i < size; i++ {
		if arr[i] == value {
			fmt.Println("Индекс первого вхождения искомого значения:", i)
			k = i
			break
		}
	}
	if k == -1 {
		fmt.Println("Искомого значения нет в массиве")
	}
}

func main() {
	value := rand.Intn(3) //такое мелкое значение, чтобы было проще затестить
	//value := 5
	fmt.Println("Искомое значение:", value)
	arr := create()
	//arr := [12]int{1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5}
	findValue(arr, value)
}
