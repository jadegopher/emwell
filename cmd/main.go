package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	termChan := make(chan os.Signal, 1)
	wg := sync.WaitGroup{}
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	ctx, cancelFunc := context.WithCancel(context.Background())

	c := &container{}
	tg, err := c.InitTelegramBot(&wg)
	if err != nil {
		panic(err)
	}

	go func(ctx context.Context, wg *sync.WaitGroup) {
		if err = tg.HandleUpdates(ctx); err != nil {
			panic(err)
		}
	}(ctx, &wg)

	for range termChan {
		cancelFunc()
	}

	wg.Wait()
}
