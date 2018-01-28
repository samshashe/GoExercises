package main

import "fmt"

func main4() {
	ar := [3]*int{0: new(int), 1: new(int), 2: new(int)}
	*ar[1] = 3
	var ar2 [2]int
	fmt.Println("ar2: ", ar2)

	var array [5][5]int
	// Set integer values to each individual element.

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			array[i][j] = (i + 1) * (j + 1)
		}
	}

	fmt.Println("array: ", array)

	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := slice1[1:3]
	slice2[0] = 4
	fmt.Println("slice1: ", slice1, "slice2", slice2)
	slice3 := []int{3: 8}
	fmt.Println("slice3: ", slice3)
	slice4 := make([]int, 3, 5)
	fmt.Println("slice4: ", slice4)

	slice5 := append(slice2, 60)
	fmt.Println("slice1: ", slice1, "slice2: ", slice2, "slice5: ", slice5)

}
