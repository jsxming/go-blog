package main

import (
	"context"
	"fmt"
	"time"
)



func worker(ctx context.Context){
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

func worker2(ctx context.Context)  {
LOOP:
	for {
		fmt.Println("worker-222")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

//func main() {
//	ctx,cancel:=context.WithCancel(context.Background())
//	go worker(ctx)
//	time.Sleep(time.Second*3)
//	cancel()// 通知子goroutine结束
//	fmt.Println("over")
//}

func main() {
	d := time.Now().Add(1 * time.Second)
	//当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭
	ctx, cancel := context.WithDeadline(context.Background(), d)
	//context.WithTimeout() 本质上等价于 context.WithDeadline

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("---",ctx.Err())
	}

}