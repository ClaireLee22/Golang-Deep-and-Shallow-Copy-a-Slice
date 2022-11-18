package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	copy(dst, src)

	fmt.Printf("src slice: %v, memory address of the backing array: %p\n", src, src)
	fmt.Printf("dst slice: %v, memory address of the backing array: %p\n", dst, dst)
}

/*
output:
src slice: [1 2 3 4 5], memory address of the backing array: 0xc0000aa030
dst slice: [1 2 3 4 5], memory address of the backing array: 0xc0000aa060
*/
