package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/payfazz/fazzlearning-api/config"
	"github.com/payfazz/fazzlearning-api/http/route"
)

// APIServer ...
type APIServer struct{}

// Serve is a function for stating API Server
func (as *APIServer) Serve() {
	s := &http.Server{
		Addr:         config.Get(config.PORT),
		ReadTimeout:  config.GetIfDuration(config.I_READ_TIMEOUT),
		WriteTimeout: config.GetIfDuration(config.I_WRITE_TIMEOUT),
		IdleTimeout:  config.GetIfDuration(config.I_IDLE_TIMEOUT),
		Handler:      route.Compile(),
	}

	serverErrCh := make(chan error)
	go func() {
		defer close(serverErrCh)
		log.Println(fmt.Sprintf("Server running at port %v", config.Get(config.PORT)))
		serverErrCh <- s.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signals := []os.Signal{
		syscall.SIGTERM,
		syscall.SIGINT,
	}
	signal.Notify(signalChan, signals...)

	select {
	case err := <-serverErrCh:
		log.Println("Server returning error: ", err)
	case sig := <-signalChan:
		signal.Reset(signals...)
		waitFor := config.GetIfDuration(config.I_WAIT_SHUTDOWN)

		log.Printf("Got '%s' signal, Stopping (Waiting for graceful shutdown: %s)\n", sig.String(), waitFor.String())

		ctx, cancel := context.WithTimeout(context.Background(), waitFor)
		defer cancel()

		if nil != s.Shutdown(ctx) {
			log.Println("Shutting down server returning error", s.Shutdown(ctx))
		} else {
			log.Println("Shutting down server")
		}
	}
}

// CreateAPIServer is a function to create an API Server
func CreateAPIServer() *APIServer {
	return &APIServer{}
}
