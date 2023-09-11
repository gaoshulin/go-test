package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// 爬取邮箱正则
	reEmail = `\w+@\w+\.\w+`
	// 爬取网址
	reLinke = `href="(https?://[\s\S]+?)"`
	// 爬取手机号码
	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	// 爬取身份证号码
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	// 爬取图片
	reImg = `http?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	// 1.抽取的爬邮箱
	var url = "https://tieba.baidu.com/p/8345823124"
	GetEmail(url)

	// 2.爬链接
	GetLink(url)

	// 3.爬手机号
	GetPhone(url)

	// 4.爬身份证号
	GetIdCard(url)

	// 5.爬图片
	GetImg(url)
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "Http.Get url")
	defer resp.Body.Close()

	// 读取网页内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")

	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func GetEmail(url string) {
	pageStr := GetPageStr(url)
	res := regexp.MustCompile(reEmail)
	results := res.FindAllStringSubmatch(pageStr, -1)

	fmt.Println("=============== Email ===============")
	for _, r := range results {
		fmt.Println(r)
	}
}

func GetLink(url string) {
	pageStr := GetPageStr(url)
	res := regexp.MustCompile(reLinke)
	results := res.FindAllStringSubmatch(pageStr, -1)

	fmt.Println("=============== Link ===============")
	for _, result := range results {
		fmt.Println(result[1])
	}
}

func GetPhone(url string) {
	pageStr := GetPageStr(url)
	res := regexp.MustCompile(rePhone)
	results := res.FindAllStringSubmatch(pageStr, -1)

	fmt.Println("=============== Phone ===============")
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	res := regexp.MustCompile(reIdcard)
	results := res.FindAllStringSubmatch(pageStr, -1)

	fmt.Println("=============== IdCard ===============")
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	res := regexp.MustCompile(reImg)
	results := res.FindAllStringSubmatch(pageStr, -1)

	fmt.Println("=============== Image ===============")
	for _, result := range results {
		fmt.Println(result[0])
	}
}
