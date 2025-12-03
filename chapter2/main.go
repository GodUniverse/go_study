package main

import (
	"fmt"
	"time"
)

func addTen(p *int) {
	*p += 10
}

func doubleSlide(p *[]int) {
	for i, v := range *p {
		(*p)[i] = v * 2
	}
}

func printOddNumbers() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Print(i, " ")
		}
	}
}

func printEvenNumbers() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}

func taskScheduler(tasks []func()) {
	for i, task := range tasks {
		go func(taskId int, t func()) {
			startTime := time.Now()
			t()
			fmt.Printf("Task %v took %v\n", t, time.Since(startTime))
		}(i, task)
	}
}

func main() {
	//指针题目一
	a := 10
	addTen(&a)
	fmt.Println(a)

	//指针题目二
	slice := []int{1, 2, 3, 4, 5}
	doubleSlide(&slice)
	fmt.Println(slice)

	//协程题目一
	go printOddNumbers()
	go printEvenNumbers()
	time.Sleep(1 * time.Second)

	//协程题目二
	tasks := []func(){
		func() { time.Sleep(2 * time.Second); fmt.Println("任务1完成") },
		func() { time.Sleep(1 * time.Second); fmt.Println("任务2完成") },
		func() { time.Sleep(3 * time.Second); fmt.Println("任务3完成") },
	}
	taskScheduler(tasks)
	time.Sleep(4 * time.Second)
}
