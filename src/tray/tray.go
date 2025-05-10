//go:build darwin || windows

package tray

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/sech1p/AppleMusic_Discord_RPC/src/config"
	"github.com/skratchdot/open-golang/open"
)

var Status string

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("🎧")
	mInfo := systray.AddMenuItem(fmt.Sprintf("Apple Music Presence %s by sech1p ✨", config.Version), "Visit project website")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Close this amazing Presence 😿")

	// because I wasn't able to force the click condition
	// and title setting to work in one condition
	go func() {
		select {
		case <-mInfo.ClickedCh:
			open.Run("https://github.com/sech1p/AppleMusic_Discord_RPC")
		case <-mQuit.ClickedCh:
			systray.Quit()
		}
	}()

	for {
		systray.SetTitle("🎧‍" + Status)
	}
}

func onExit() {
	os.Exit(0)
}

func Init() {
	systray.Run(onReady, onExit)
}
