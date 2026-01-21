package thread_pool

import (
	"context"
	"sync"
)

type TPContract interface {
	CreateTask(task func())
	Work(ctx context.Context, num int, wg *sync.WaitGroup)
}
