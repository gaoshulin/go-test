package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userinfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userinfo)

	// map的遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

}
