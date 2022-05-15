// package slicetest
package main

import "fmt"

func update_slice(arr []int) {
	arr[0] = 999
	m1 := map[int]string{1: "a", 2: "b"}
	if val, ok := m1[10]; !ok {
		val = "ok"
		fmt.Printf("val=%s\n", val)
	} else {
		fmt.Printf("already exist!, val=%s\n", val)
	}
	fmt.Println("m1=", m1)
	fmt.Printf("arr before append %p, %d, %d\n", arr, len(arr), cap(arr))
	arr = append(arr, 2, 3, 1)
	fmt.Printf("arr after append %p, %d, %d\n", arr, len(arr), cap(arr))
}

func change_ele_by_pt(arr *[]int, i int) {
	(*arr)[i] = 666
	fmt.Printf("change_ele_by_pt %p, %d, %d\n", *arr, len(*arr), cap(*arr))
}

func change_ele_by_arr(arr []int, i int) {
	arr[i] = 666
	fmt.Printf("change_ele_by_arr %p, %d, %d\n", arr, len(arr), cap(arr))
}

func main() {
	crr := []int{99, 88, 77, 66, 55}
	update_slice(crr)
	fmt.Printf("crr in main before change %p, %d, %d\n", crr, len(crr), cap(crr))
	fmt.Println(crr)
	change_ele_by_pt(&crr, 0)
	fmt.Printf("crr in main after change by pt %p, %d, %d\n", crr, len(crr), cap(crr))
	fmt.Println(crr)
	change_ele_by_arr(crr, 0)
	fmt.Printf("crr in main after change by arr %p, %d, %d\n", crr, len(crr), cap(crr))
	fmt.Println(crr)
}
