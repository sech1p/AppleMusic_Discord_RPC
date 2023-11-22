//go:build darwin || windows

package main

import (
	"github.com/sech1p/AppleMusic_Discord_RPC/src/parser"
	"github.com/sech1p/AppleMusic_Discord_RPC/src/platform"
	"github.com/sech1p/AppleMusic_Discord_RPC/src/presence"
	"github.com/sech1p/AppleMusic_Discord_RPC/src/tray"
)

func main() {
	err := presence.Login()
	if err != nil {
		tray.Status = "🔑🚫"
	}

	for {
		go func() {
			for {
				out, err := platform.GetSongData()
				if err != nil {
					switch err.Error() {
					case "no song is playing":
						tray.Status = "‍⏳"
					}
				} else {
					song := parser.Parse(out)

					err := presence.Update(song)
					if err != nil {
						tray.Status = "🚫"
					}

					tray.Status = ""
				}
			}
		}()

		tray.Init()
	}
}
