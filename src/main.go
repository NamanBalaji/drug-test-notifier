package main

import (
	"fmt"

	"github.com/go-rod/rod/lib/launcher"

	"github.com/NamanBalaji/drug-test-notifier/pkg/browser_automation"
	"github.com/NamanBalaji/drug-test-notifier/pkg/config"
	"github.com/NamanBalaji/drug-test-notifier/pkg/data"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	browser, l := browser_automation.GetBrowser(false)
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

	fmt.Println(d)
}
