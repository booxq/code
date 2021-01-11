package main

import "fmt"

func foo(slice []int) {
	slice[0] = 1000
}

func main() {

	var arr = [4]int{1, 2, 3, 4}
	s1 := arr[0:4]
	s2 := s1[2:3:3]
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	fmt.Println("slice1 cap", cap(s1))
	fmt.Println("slice2 cap", cap(s2))
	fmt.Println("slice1 addr=", &s1[2])
	fmt.Println("slice2 addr=", &s2[0])
	//s2 = append(s2, 5)
	foo(s2)
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	fmt.Println("slice1 cap", cap(s1))
	fmt.Println("slice2 cap", cap(s2))
	fmt.Println("slice1 addr=", &s1[2])
	fmt.Println("slice2 addr=", &s2[0])

	/*二维切片*/

	// slice := [][]int{{10, 20, 30}, {100, 200, 300}}
	// fmt.Println("slice", &slice)
	// fmt.Println("00 addr", &slice[0][0])
	// fmt.Println("10 addr", &slice[1][0])
	// fmt.Println("11 addr", &slice[1][1])
	// slice[0] = append(slice[0], 20)
	// fmt.Println("00 addr", &slice[0][0])

}
