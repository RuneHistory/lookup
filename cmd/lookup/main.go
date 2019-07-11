package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"log"
	"lookup/internal/application/service"
	"lookup/internal/transport/http_transport"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	address := os.Getenv("LISTEN_ADDRESS")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	shutdownCh := make(chan struct{})
	errCh := make(chan error)
	go handleShutdownSignal(shutdownCh)
	go func() {
		select {
		case <-shutdownCh:
			break
		case err := <-errCh:
			log.Printf("fatal error: %s", err)
		}
		cancel()
	}()

	r := chi.NewRouter()

	highScoreService := service.NewHighScoreService()
	http_transport.Bootstrap(r, highScoreService)

	wg.Add(1)
	go http_transport.Start(address, r, wg, ctx, errCh)

	// doneCh will be closed once wg is done
	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-doneCh:
		// we're finished so start the shutdown
		log.Println("all services finished")
	case <-ctx.Done():
		break
		// break out and wait for shutdown
	}

	log.Println("waiting for shutdown")

	select {
	case <-time.After(time.Second * 10):
		log.Println("killed - took too long to shutdown")
	case <-doneCh:
		log.Println("all services shutdown")
	}
}

func handleShutdownSignal(shutdownCh chan struct{}) {
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	startedShutdown := false
	for {
		<-quitCh
		if startedShutdown {
			os.Exit(0)
		}
		close(shutdownCh)
		startedShutdown = true
	}
}
