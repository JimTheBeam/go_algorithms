package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{-741, 123, 369, -987, 753, -5, 159, -852, -2, 35, -94, 72, -85, 12, 34, -97, 65}
	fmt.Println("len(arr):", len(arr))
	pivot := rand.Intn(24)
	fmt.Printf("pivot: %d\n", pivot)
	// fmt.Println(sorted)
}
