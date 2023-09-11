package main

import "fmt"

func main() {
	// 指针取值
	a := 10
	b := &a

	// 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	fmt.Printf("value of b:%v\n", b)

	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是:%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}

}
