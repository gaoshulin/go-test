package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Userinfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	// 单独写回调函数
	http.HandleFunc("/index", myHandle)

	// 解析 template 模版
	http.HandleFunc("/hello", sayHello)

	http.HandleFunc("/login", login)

	// 监听地址
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func myHandle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr, "连接成功")

	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", request.Method)
	fmt.Println("url:", request.URL.Path)
	fmt.Println("header:", request.Header)
	fmt.Println("body:", request.Body)

	// 回复
	_, err := writer.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func sayHello(write http.ResponseWriter, request *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./tmpl/http.tmpl", "./tmpl/http-ul.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := Userinfo{
		"galen",
		"male",
		26,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(write, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./tmpl/login.tmpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		err := r.ParseForm()
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}

		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("email:", r.Form["email"])
	}
}
