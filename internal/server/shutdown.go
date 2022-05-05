package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// clean up operation with server shutdown
type Operation func() error

func GracefulShutdown(ctxBase context.Context, timeout time.Duration, ops map[string]Operation) (ctx context.Context, wait chan struct{}) {
	wait = make(chan struct{})
	ctx, cancel := context.WithCancel(ctxBase)
	go func() {
		// use buffered channel to prevent memory leak
		sc := make(chan os.Signal, 1)

		// check any syscalls that shutdown the server
		signal.Notify(sc, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
		// blocking
		<-sc

		log.Println("start shutting down process")

		// set timeout for all the operations to prevent hung
		timeoutF := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force system exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutF.Stop()

		var wg sync.WaitGroup

		// process shutdown operations async
		for k, o := range ops {
			wg.Add(1)
			key := k
			opr := o
			go func() {
				defer wg.Done()
				log.Printf("cleanning up operation: '%s'", key)
				if err := opr(); err != nil {
					log.Printf("'%s' operation failed with error: '%+v", key, err)
					return
				}
				log.Printf("cleanning up operation successfully: '%s'", key)
			}()
		}
		wg.Wait()
		cancel()
		close(wait)
	}()
	return ctx, wait
}

// reference
// https://habibfikri.medium.com/graceful-shutdown-in-go-548fd7d8094c
