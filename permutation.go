package main

import "fmt"

func Permutation(data *[]int, i int, length int) {
	if length == i {
		PrintSlice(data)
		return
	}
	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func swap(data *[]int, x int, y int) {
	(*data)[x], (*data)[y] = (*data)[y], (*data)[x]
}

func main() {
	data := make([]int, 5)
	for i := 0; i < 5; i++ {
		data[i] = i
	}
	Permutation(&data, 0, 5)
}

func PrintSlice(data *[]int) {
	for i := 0; i < len(*data); i++ {
		fmt.Print((*data)[i], " ")
	}
	fmt.Println()
}
