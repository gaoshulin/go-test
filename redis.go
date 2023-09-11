package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed.", err)
		return
	}
	fmt.Println("redis conn success")

	defer con.Close()

	// 1、string Set
	//_, err = con.Do("Set", "later", 100)
	//// 设置过期时间
	//con.Do("expire", "later", 30)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//// 获取 Get
	//r, err := redis.Int(con.Do("Get", "later"))
	//if err != nil {
	//	fmt.Println("get abc failed,", err)
	//	return
	//}
	//fmt.Println(r)

	// 2、list
	//_, err = con.Do("lpush", "list", "later", "galen", "alan")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//res, err := redis.String(con.Do("lpop", "list"))
	//if err != nil {
	//	fmt.Println("get list failed.", err)
	//	return
	//}
	//fmt.Println(res)

	// 3、hash
	_, err = con.Do("HSet", "books", "test", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := redis.Int(con.Do("HGet", "books", "test"))
	if err != nil {
		fmt.Println("get books failed.", err)
		return
	}
	fmt.Println(res)
}
