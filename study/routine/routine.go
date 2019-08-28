package routine

import (
	"fmt"
	"runtime"
	"sync"
)

func Sync_test() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:= ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:== ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type student struct {
	Name string
	Age  int
}

//下面函数存在的问题是：使用地址给数组元素赋值，导致所有数组元素的值相同。
func Pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println("m: ", m)
}

//defer的执行是一个入栈出栈的过程
func Defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
