package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	systray.SetIcon(icon)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		addCheckItem(text)
	}

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "")
	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}
}

func addCheckItem(label string) {
	mChecked := systray.AddMenuItem(label, "")
	go func() {
		for {
			select {
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
				} else {
					mChecked.Check()
				}
			}
		}
	}()
}
