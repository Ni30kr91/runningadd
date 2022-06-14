package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	var final = []int{}
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]

		final = append(final, sum)

	}

	fmt.Println(final)
}
