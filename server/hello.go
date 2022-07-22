package main

import (
	"coolcar/constraints"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("sum base on generics:", sum([]float64{0.1, 0.2, 0.3}))
}

func sum[T constraints.Ordered](nums []T) (ans T) {
	for _, num := range nums {
		ans += num
	}
	return ans
}
