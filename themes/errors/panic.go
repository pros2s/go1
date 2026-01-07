package errors

import "fmt"

func PanicTest(index int) {
	defer func() {
		p := recover()

		if p != nil {
			fmt.Println(p)
		}
	}()

	array := []int{1, 2, 3, 4}
	fmt.Println(array[index])
}
