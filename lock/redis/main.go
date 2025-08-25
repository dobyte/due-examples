package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dobyte/due/lock/redis/v2"
	"github.com/dobyte/due/v2/lock"
)

func main() {
	// 设置锁制造商
	lock.SetMaker(redis.NewMaker())

	var (
		wg    sync.WaitGroup
		ctx   = context.Background()
		total int
	)

	wg.Add(100)

	startTime := time.Now().UnixNano()

	for range 100 {
		go func() {
			locker := lock.Make("lock")

			if err := locker.Acquire(ctx); err != nil {
				return
			}

			defer locker.Release(ctx)

			total++

			wg.Done()
		}()
	}

	wg.Wait()

	totalTime := float64(time.Now().UnixNano()-startTime) / float64(time.Second)

	fmt.Printf("total		: %d\n", total)
	fmt.Printf("latency		: %fs\n", totalTime)
}
