package main

import (
	"fmt"
)

func main() {
	arrA := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	arrB := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	fmt.Println("3 -> ")

	arrA2 := [][]int{
		{0, 1},
		{1, 1},
	}
	arrB2 := [][]int{
		{1, 1},
		{1, 0},
	}
	fmt.Println("2 -> ")
}
