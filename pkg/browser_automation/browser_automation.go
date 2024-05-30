package browser_automation

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"github.com/NamanBalaji/drug-test-notifier/pkg/config"
	"github.com/NamanBalaji/drug-test-notifier/pkg/consts"
	"github.com/NamanBalaji/drug-test-notifier/pkg/data"
)

// function to log in
// function to open page
// function to fill in the login form

func GetBrowser(headless bool) (*rod.Browser, *launcher.Launcher) {
	if headless {
		return rod.New(), nil
	}

	l := launcher.New().Headless(false)

	url := l.MustLaunch()
	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(2 * time.Second)

	return browser, l
}

func login(cfg config.Config, browser *rod.Browser) {
	page := browser.MustPage(consts.LoginPageURL).MustWaitStable()
	defer page.MustClose()

	page.MustElement(consts.ProgramIdSelector).MustInput(cfg.ProgramId)
	page.MustElement(consts.UsernameSelector).MustInput(cfg.Username)
	page.MustElement(consts.PasswordSelector).MustInput(cfg.Password)

	page.MustElement(consts.LoginButtonSelector).MustClick()

	page.Race().Element("#loginBox").MustHandle(func(_ *rod.Element) {
		panic("Cannot login")
	}).Element("#bd_b").MustHandle(func(_ *rod.Element) {
		log.Println("logged in successfully")
	}).MustDo()

}

func goToIFrameAndClickTestStatus(browser *rod.Browser, data *data.Data) error {
	page := browser.MustPage(consts.IFrameURL).MustWaitStable()
	// get bills due
	bills, err := getBillsDue(page)
	if err != nil {
		return err
	}
	data.BillsDue = bills

	//
	return nil
}

func getBillsDue(page *rod.Page) (int, error) {
	text := page.MustElement(consts.BillsSelector).MustText()
	numberString := strings.Split(text, " ")

	bills, err := strconv.Atoi(numberString[0])
	if err != nil {
		return 0, errors.New("cannot convert bills to number")
	}

	return bills, nil
}

func clickTestStatus(page *rod.Page) {
	page.MustElement(consts.TestStatusSelector).MustClick()
}
