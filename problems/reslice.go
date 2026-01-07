package problems

import "fmt"

func mutateArray(arr []int) []int {
	arr[0] = 15
	arr[1] = 10

	return append(arr, 20, 30)
}

func TestReslice() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := arr[1:3:5]

	arr2 = mutateArray(arr2)

	fmt.Println(arr)
	fmt.Println(arr2)
}
