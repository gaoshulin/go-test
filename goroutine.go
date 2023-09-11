package main

import (
	"fmt"
	"math/rand"
	//"sync"
)

//var wg sync.WaitGroup

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func main() {
	for i := 0; i < 10; i++ {
		// 启动一个goroutine就登记+1
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()

	// ===== 进程池 ==========
	// job 管道
	jobChan := make(chan *Job, 128)
	// 结果管道
	resultChan := make(chan *Result, 128)
	// 创建工作池
	createPool(64, jobChan, resultChan)
	// 打印协程数据
	go func(resultCh chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

func hello(i int) {
	defer wg.Done()

	fmt.Println("Hello Goroutine!", i)
}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
