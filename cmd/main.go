package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"emwell/internal/api/http"
)

//go:generate  protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:./internal/api/protobuf --swagger_out=allow_merge=true,merge_file_name=api:. --go_out=plugins=grpc:./internal/api/protobuf ./internal/api/protobuf/api.proto

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

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		if err = tg.HandleUpdates(ctx); err != nil {
			panic(err)
		}
	}(ctx, &wg)

	srv := http.NewServer(c.services.logger, ":17002", c.InitHttpHandlers())

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		if err = srv.ListenAndServe(ctx); err != nil {
			panic(err)
		}
	}(ctx, &wg)

	for range termChan {
		cancelFunc()
		break
	}

	wg.Wait()
}
