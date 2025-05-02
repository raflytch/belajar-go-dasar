package main

import "fmt"

func main() {
	numbers := make([]int, 0, 5)

	for i := 1; i <= 5; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println("Slice", numbers)
	fmt.Println("Panjang Slice", len(numbers))
	fmt.Println("Kapasitas Slice", cap(numbers))
}
