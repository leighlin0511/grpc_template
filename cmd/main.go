package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leighlin0511/grpc_template/internal/app"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		osCall := <-c
		log.Printf("system call: %+v", osCall)
		cancel()
	}()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
