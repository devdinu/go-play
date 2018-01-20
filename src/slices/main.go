package main

import "fmt"

func main() {
	//	bla := [4]int{1, 2, 3, 4}
	//	//x := [3]int{1, 3, 4}
	//	mutate(bla)
	//	fmt.Printf("%+v, %p\n", bla, &bla)

	//	sl := []int{1, 2, 3, 4}
	//	mutateSlice(&sl)

	//slices(11)
	//sl := []int{1, 2, 3, 4}
	//var sl []int
	bla := make([]int, 3, 8)
	smaller := bla[1:3]
	//fmt.Println("asdfasf:", bla, sl, bla[6])
	fmt.Printf("c:%d, l:%d", cap(smaller), len(smaller))

	//bla[7]
	//	fmt.Printf("main: %+v, %p\n", sl, &sl)
	//	mutate(sl)
	//	//slicing(sl)
	//	fmt.Printf("main: %+v, %p\n", sl, &sl)
}

func mutate(data []int) {
	data[0] = 100
	fmt.Println(data, &data)
}

func mutateSlice(data *[]int) {

}
