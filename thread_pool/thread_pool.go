package thread_pool

import (
	"context"
	"fmt"
	"sync"
)

type ThreadPool struct {
	Threads   *sync.WaitGroup
	TaskQueue chan func()
	ctx       context.Context
}

func NewThreadPool(ctx context.Context, workers_number int) *ThreadPool {
	pool := &ThreadPool{
		TaskQueue: make(chan func(), 10000),
		ctx:       ctx,
	}

	wg := &sync.WaitGroup{}
	for i := range workers_number {
		wg.Add(1)
		go pool.Work(ctx, i, wg)
	}
	pool.Threads = wg

	return pool
}

func (tp *ThreadPool) CreateTask(task func()) {
	select {
	case tp.TaskQueue <- task:
		// send task to queue
	default:
		// unable to send task so skip it
	}
}

func (tp *ThreadPool) Work(ctx context.Context, num int, wg *sync.WaitGroup) {
	fmt.Printf("Started goroutine %d\n", num)
	defer wg.Done()
	for {
		select {
		case task, ok := <-tp.TaskQueue:
			if !ok {
				return
			}
			fmt.Printf("Goroutine %d working\n", num)
			task()
		case <-ctx.Done():
			fmt.Println("return")
			return
		}
	}
}

var _ TPContract = (*ThreadPool)(nil)
