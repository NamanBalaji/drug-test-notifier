package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-rod/rod/lib/launcher"

	"github.com/NamanBalaji/drug-test-notifier/pkg/browser_automation"
	"github.com/NamanBalaji/drug-test-notifier/pkg/config"
	"github.com/NamanBalaji/drug-test-notifier/pkg/data"
	"github.com/NamanBalaji/drug-test-notifier/pkg/mail"
	"github.com/NamanBalaji/drug-test-notifier/pkg/server"
)

func main() {
	cfg := config.LoadConfig()
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	headless := false
	if !*debug {
		headless = true
	}
	triggerChan := make(chan server.Trigger)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		automate(cfg, triggerChan, done, headless)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.RunServer(cfg, triggerChan, done); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not start server: %v\n", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("Received shutdown signal")

	// Signal the server and run function to stop
	close(done)

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Server and processing gracefully stopped")
}

func automate(cfg config.Config, triggerChan chan server.Trigger, done chan struct{}, headless bool) {
	browser, l := browser_automation.GetBrowser(headless)
	if l != nil {
		defer l.Cleanup()

		browser.MustConnect()
		launcher.Open(browser.ServeMonitor(""))
	} else {
		browser.MustConnect()
	}
	defer browser.MustClose()

	d := data.NewData()

	for {
		select {
		case <-triggerChan:
			fmt.Println("Starting automation...")
			err := browser_automation.Run(cfg, browser, d)
			if err != nil {
				fmt.Println(err)
			}

			mailClient := mail.NewMailClient(cfg.SenderEmail, cfg.AppPassword)

			subject, body, err := d.GenerateEmail()
			if err != nil {
				log.Println("Error while browser automation: ", err)

				continue
			}

			err = mailClient.SendMail(subject, body, cfg.Username)
			if err != nil {
				log.Println("Error while sending email: ", err)
			}

			fmt.Println("Email sent successfully!")
		case <-done:
			fmt.Println("Stopping processing")
			return
		}
	}
}
