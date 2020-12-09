package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("start")
	g := errgroup.Group{}
	httpServerErr := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)

	s := http.Server{Addr: ":8080"}

	g.Go(func() error {
		for {
			httpServerErr <- s.ListenAndServe()
			break
		}
		select {
		case err := <-httpServerErr:
			fmt.Println("连接")
			close(sigChan)
			close(httpServerErr)
			return err
		}
	})

	g.Go(func() error {
		signal.Notify(sigChan, syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-sigChan
		fmt.Println("linux 信号")
		return s.Shutdown(context.TODO())
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
