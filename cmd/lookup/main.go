package main

import (
	"errors"
	"github.com/go-chi/chi"
	"log"
	"lookup/internal/application/service"
	"lookup/internal/transport/http_transport"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	address := os.Getenv("LISTEN_ADDRESS")

	wg := &sync.WaitGroup{}
	shutdownCh := make(chan struct{})
	errCh := make(chan error)
	go handleShutdownSignal(errCh)

	r := chi.NewRouter()

	highScoreService := service.NewHighScoreService()
	http_transport.Bootstrap(r, highScoreService)

	go http_transport.Start(address, r, wg, shutdownCh, errCh)

	err := <-errCh
	if err != nil {
		log.Printf("fatal err: %s\n", err)
	}

	log.Println("initiating graceful shutdown")
	close(shutdownCh)

	wg.Wait()
	log.Println("shutdown")
}

func handleShutdownSignal(errCh chan error) {
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	hit := false
	for {
		<-quitCh
		if hit {
			os.Exit(0)
		}
		if !hit {
			errCh <- errors.New("shutdown signal received")
		}
		hit = true
	}
}
