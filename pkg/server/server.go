package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NamanBalaji/drug-test-notifier/pkg/config"
)

type Trigger struct{}

func RunServer(cfg config.Config, triggerChan chan Trigger, done chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/trigger", func(w http.ResponseWriter, r *http.Request) {
		triggerChan <- Trigger{}
		fmt.Println("Sent trigger")
		w.WriteHeader(http.StatusOK)

		return
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	// Run the server in a separate goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	fmt.Println("Server started on :8080")

	<-done // Wait for the done signal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	fmt.Println("Server stopped")

	return nil
}
