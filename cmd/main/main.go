package main

import (
	"context"
	"fmt"
	"go-thread-pool/thread_pool"
)

func main() {
	thread_count := 4

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp := thread_pool.NewThreadPool(ctx, thread_count)

	for i := range 100 {
		num := i
		tp.CreateTask(func() { fmt.Printf("Message %d\n", num) })
	}

	close(tp.TaskQueue)
	tp.Threads.Wait()
}
