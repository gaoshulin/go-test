package main

import "fmt"

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

var arr01 [5][3]int
var arr02 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

// 数组
func main() {
	// 一维数组
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)

	// 多维数组
	a1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b1 := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr01, arr02)
	fmt.Println(a1, b1)
}
