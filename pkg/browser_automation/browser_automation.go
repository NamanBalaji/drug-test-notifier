package browser_automation

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"github.com/NamanBalaji/drug-test-notifier/pkg/config"
	"github.com/NamanBalaji/drug-test-notifier/pkg/data"
)

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
	page := browser.MustPage(loginPageURL).MustWaitStable()
	defer page.MustClose()

	page.MustElement(programIdSelector).MustInput(cfg.ProgramId)
	page.MustElement(usernameSelector).MustInput(cfg.Username)
	page.MustElement(passwordSelector).MustInput(cfg.Password)

	page.MustElement(loginButtonSelector).MustClick()

	page.Race().Element("#loginBox").MustHandle(func(_ *rod.Element) {
		panic("Cannot login")
	}).Element("#bd_b").MustHandle(func(_ *rod.Element) {
		log.Println("logged in successfully")
	}).MustDo()

}

func goToIFrameAndClickTestStatus(browser *rod.Browser, data *data.Data) error {
	page := browser.MustPage(iFrameURL).MustWaitStable()
	defer page.MustClose()

	// get bills due
	bills, err := getBillsDue(page)
	if err != nil {
		return err
	}
	data.BillsDue = bills

	// click the test status button
	page.MustElement(testStatusButtonSelector).MustClick()
	err = getSelectedInfo(page, data)
	if err != nil {
		return err
	}

	return nil
}

func getSelectedInfo(page *rod.Page, data *data.Data) error {
	var err error
	page.Race().Element(confirmationNumberSpanSelector).MustHandle(func(e *rod.Element) {
		confirmation := e.MustText()
		confirmationNumber, err := strconv.Atoi(strings.TrimSpace(confirmation))
		if err != nil {
			err = fmt.Errorf("error converting confirmation number to string: %s", confirmation)
		}
		data.ConfirmationNumber = confirmationNumber
	}).MustDo()

	selectionMessage := page.MustElement(selectionStatusSelectorYes).MustText()
	if selectionMessage == "" {
		selectionMessage = page.MustElement(selectionStatusSelectorNo).MustText()
	}

	selectionMessageBool := strings.Split(selectionMessage, " - ")
	yesOrNo := strings.Split(selectionMessageBool[0], " / ")
	if yesOrNo[0] == "YES" {
		data.Selected = true
	}

	dateParse := strings.Split(selectionMessage, "(")
	dateDirty := dateParse[1]
	endIndex := strings.Index(dateDirty, ")")
	cleanedDateString := dateDirty[:endIndex]
	data.Date = cleanedDateString

	return err
}

func getBillsDue(page *rod.Page) (int, error) {
	text := page.MustElement(billsSelector).MustText()
	numberString := strings.Split(text, " ")

	bills, err := strconv.Atoi(numberString[0])
	if err != nil {
		return 0, errors.New("cannot convert bills to number")
	}

	return bills, nil
}

func logout(browser *rod.Browser) {
	page := browser.MustPage(logoutPage).MustWaitStable()
	defer page.MustClose()
}

func Run(cfg config.Config, browser *rod.Browser, data *data.Data) error {
	login(cfg, browser)

	err := goToIFrameAndClickTestStatus(browser, data)
	if err != nil {
		return err
	}

	logout(browser)

	return nil
}
