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
	ctx, stop := signal.NotifyContext(context.Background(), []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	defer stop()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
