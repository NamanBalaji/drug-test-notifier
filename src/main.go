package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-rod/rod/lib/launcher"

	"github.com/NamanBalaji/drug-test-notifier/pkg/browser_automation"
	"github.com/NamanBalaji/drug-test-notifier/pkg/config"
	"github.com/NamanBalaji/drug-test-notifier/pkg/data"
	"github.com/NamanBalaji/drug-test-notifier/pkg/mail"
)

func main() {
	cfg := config.LoadConfig()
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	headless := false
	if !*debug {
		headless = true
	}

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
	err := browser_automation.Run(cfg, browser, d)
	if err != nil {
		fmt.Println(err)
	}

	mailClient := mail.NewMailClient(cfg.SenderEmail, cfg.AppPassword)

	subject, body, err := d.GenerateEmail()
	if err != nil {
		log.Fatal(err)
	}

	err = mailClient.SendMail(subject, body, cfg.Username)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}
