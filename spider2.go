package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	// 用于监控协程
	chanTask chan string

	waitGroup sync.WaitGroup
	reImage   = `http?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func HandleErrors(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 爬取图片
func main() {
	// 定义需要爬取的网址，失效了自定更换地址
	var url = "https://tieba.baidu.com/p/8345823124"

	// 1、初始化管道
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 26)

	// 2、爬虫协程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go GetImgUrls(url)
	}

	// 3、任务统计协程，统计26个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()

	// 4、下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}

func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++

		if count == 26 {
			close(chanImageUrls)
			break
		}
	}

	waitGroup.Done()
}

// 爬图片链接到管道
func GetImgUrls(url string) {
	urls := getImages(url)

	// 遍历切片李所有的链接，并存入管道
	for _, url := range urls {
		chanImageUrls <- url
	}

	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	waitGroup.Done()
}

// 获取当前页图片链接
func getImages(url string) (urls []string) {
	pageStr := GetPageStr2(url)
	res := regexp.MustCompile(reImage)
	results := res.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))

	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

// 抽取根据url获取内容
func GetPageStr2(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleErrors(err, "http.Get url")
	defer resp.Body.Close()

	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleErrors(err, "ioutil.ReadAll")

	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

// 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// 下载图片，传入的是图片地址
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleErrors(err, "http.get.url")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HandleErrors(err, "resp.body")
	filename = "D:/Go/www/src/test/images/" + filename

	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}
